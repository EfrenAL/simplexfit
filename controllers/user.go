package controllers

import (
	"errors"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

type User struct {
	gorm.Model
	ID        string
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Weight    int       `json:"weight"`
	Height    int       `json:"hheight"`
	Workouts  []Workout `json:"workouts" gorm:"many2many:user_workout;"`
	WorkoutID uint      `json:"-"`
}

func initializeAppDefault(ctx *gin.Context) *firebase.App {
	// Get variables
	path := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	projectId := os.Getenv("PROJECT-ID")
	//Open Firebase project
	opt := option.WithCredentialsFile(path)
	config := &firebase.Config{ProjectID: projectId}
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	enableUser(ctx, app)

	return app
}

func enableUser(ctx *gin.Context, app *firebase.App) {
	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
	params := (&auth.UserToUpdate{}).
		Email("efrenalla@ggmail.com").
		EmailVerified(true).
		Disabled(false)

	client.UpdateUser(ctx, "DBW5XQODv9OYNqLBmRKdQkuZGcR2", params)
}

func GetUserController(ctx *gin.Context) {

	idToken := ctx.Request.Header.Get("Authorization")

	firebaseApp := initializeAppDefault(ctx)
	jwtToken, err := verifyIDToken(ctx, firebaseApp, idToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Invalid token",
		})
		return
	}
	user := getUserFirebase(ctx, firebaseApp, jwtToken.Subject)
	//ToDo check getUserFirebase return values
	if user == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Invalid token",
		})
		return
	}
	//Account is not verfied
	if !user.EmailVerified {
		ctx.JSON(http.StatusForbidden, gin.H{
			"status":  http.StatusForbidden,
			"message": "Account not verified",
		})
		return
	}
	//Check if account created
	userDb, err := getUserDbById(user.TenantID)

	if err != nil {
		userCreated, _ := createUser(user)
		//Create user calll
		ctx.JSON(http.StatusCreated, gin.H{
			"status":  http.StatusCreated,
			"message": "User Created Succesfully",
			"data":    userCreated,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "User Retrived Succesfully",
		"data":    userDb,
	})
}

func getUserFirebase(ctx *gin.Context, app *firebase.App, uid string) *auth.UserRecord {
	// Get an auth client from the firebase.App
	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	u, err := client.GetUser(ctx, uid)
	if err != nil {
		log.Fatalf("error getting user %s: %v\n", uid, err)
	}
	log.Printf("Successfully fetched user data: %v\n", u)
	return u
}

func verifyIDToken(ctx *gin.Context, app *firebase.App, idToken string) (*auth.Token, error) {
	// [START verify_id_token_golang]
	client, err := app.Auth(ctx)
	if err != nil {
		log.Printf("error getting Auth client: %v\n", err)
		return nil, err
	}

	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		log.Printf("error verifying ID token: %v\n", err)
		return nil, err
	}

	log.Printf("Verified ID token: %v\n", token)
	// [END verify_id_token_golang]

	return token, nil
}

func createUser(userFirebase *auth.UserRecord) (*User, error) {

	id := userFirebase.UserInfo.UID
	email := userFirebase.Email
	user := &User{ID: id, Email: email}

	result := gormDBConnect.Create(user)

	if result.Error != nil {
		log.Printf("Error while inserting new workout into db, Reason: %v\n", result.Error)
		return nil, errors.New("error creating user")
	}

	return user, nil
}

func getUserDbById(userId string) (*User, error) {
	// Get user from the database
	user := &User{}
	result := gormDBConnect.First(&user, userId)

	if result.Error != nil {
		log.Printf("Error while getting a single todo, Reason: %v\n", result.Error)
		return nil, errors.New("no user")
	}

	return user, nil
}

func UpdateUser(ctx *gin.Context) {

	var userBody User
	ctx.BindJSON(&userBody)

	user := &User{}
	userId := ctx.Param("userId")
	result := gormDBConnect.First(&user, userId)

	if result.Error != nil {
		log.Printf("Error while getting a single user, Reason: %v\n", result.Error)
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "User not found",
		})
		return
	}

	user.Weight = userBody.Weight
	user.Height = userBody.Height
	user.Name = userBody.Name

	gormDBConnect.Save(&user)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "User updated",
		"data":    user,
	})
}
