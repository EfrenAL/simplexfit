package controllers

import (
	"os"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"

)


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
	return app
}


func GetUser(ctx *gin.Context) {	

	idToken := ctx.Request.Header.Get("Authorization") //Bearer {{token}}

	log.Printf(idToken)

	firebaseApp := initializeAppDefault(ctx)
	jwtToken := verifyIDToken(ctx, firebaseApp, idToken)	
	user := getUser(ctx, firebaseApp, jwtToken.Subject)
	if user == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
	}
	 
	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "User Retrived Succesfully" ,
		"data": user,
	})	
}

func getUser(ctx *gin.Context, app *firebase.App, uid string) *auth.UserRecord {
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

func verifyIDToken(ctx *gin.Context, app *firebase.App, idToken string) *auth.Token {
	// [START verify_id_token_golang]
	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		log.Fatalf("error verifying ID token: %v\n", err)
	}

	log.Printf("Verified ID token: %v\n", token)
	// [END verify_id_token_golang]

	return token
}