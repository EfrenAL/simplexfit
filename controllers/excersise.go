package controllers

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	guuid "github.com/google/uuid"
	"gorm.io/gorm"
)


type Exercise struct {
	ID 				string  `json:"id"`
	Name 			string `json:"name"`
	Repetitions 	int64  `json:"repetitions"`
	Duration 		int64  `json:"time"`
	Complexity 		string `json:"complexity"`
	CreatedAt 		time.Time `json:"created_at"`
	UpdatedAt 		time.Time `json:"updated_at"`
}

var dbConnect *sql.DB
var gormDBConnect *gorm.DB

func InitiateDB(db *sql.DB, gormDB *gorm.DB ) {
	dbConnect = db
	gormDBConnect = gormDB
}
// Create User Table
func CreateExerciseTable() error {

	gormDBConnect.Migrator().CreateTable(&Exercise{})

	log.Printf("Exercise table created")
	return nil
}

func MainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl.html", nil)	
}


func GetAllExercise(c *gin.Context) {
	var exercise []Exercise
	
	result := gormDBConnect.Find(&exercise)
	

	if result.Error != nil {
		log.Printf("Error while getting all exercises, Reason: %v\n", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}
	message := "All Exercise %s" + string(result.RowsAffected)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": message ,
		"data": exercise,
	})	
}

func CreateExercise(c *gin.Context) {

	var exercise Exercise

	c.BindJSON(&exercise)

	name := exercise.Name
	repetitions := exercise.Repetitions
	duration := exercise.Duration
	complexity := exercise.Complexity
	id := guuid.New().String()

	result := gormDBConnect.Create(&Exercise{
		ID: id,
		Name: name,
		Repetitions: repetitions,
		Duration: duration,
		Complexity: complexity,
		CreatedAt: time.Now(),
		UpdatedAt:time.Now(),
	})

	if result.Error != nil {
		log.Printf("Error while inserting new todo into db, Reason: %v\n", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Excercise created Successfully",
	})
}

func CreateExerciseForm(c *gin.Context) {

	if err := c.Request.ParseForm(); err != nil {
		log.Printf("ParseForm() err: %v", err)
		return
	}

	name := c.Request.FormValue("name")
	// To check how to parse int to string
	//repetitions := c.Request.FormValue("repetitions")
	//duration := c.Request.FormValue("duration")
	complexity := c.Request.FormValue("complexity")
	id := guuid.New().String()

	result := gormDBConnect.Create(&Exercise{
		ID: id,
		Name: name,
		Repetitions: 100,
		Duration: 40,
		Complexity: complexity,
		CreatedAt: time.Now(),
		UpdatedAt:time.Now(),
	})

	if result.Error != nil {
		log.Printf("Error while inserting new todo into db, Reason: %v\n", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Excercise created Successfully",
	})
}



func CreateExerciseBatch(c *gin.Context) {

	var exercises []Exercise

	c.BindJSON(&exercises)

	for _, exercise := range exercises {
		exercise.ID = guuid.New().String()
	}


	result := gormDBConnect.Create(exercises)

	if result.Error != nil {
		log.Printf("Error while inserting new todo into db, Reason: %v\n", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Excercise created Successfully",
	})
}

func GetSingleExercise(c *gin.Context) {
	exerciseId := c.Param("exerciseId")
	exercise := &Exercise{ID: exerciseId}

	result := gormDBConnect.First(&exercise)
	
	if result.Error != nil {
		log.Printf("Error while getting a single todo, Reason: %v\n", result.Error)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Exercise not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Single Exercise",
		"data": exercise,
	})
}

func DeleteExercise(c *gin.Context) {
	exerciseId := c.Param("exerciseId")
	exercise := &Exercise{ID: exerciseId}

	result := gormDBConnect.Delete(exercise)

	if result.Error != nil {
		log.Printf("Error while deleting a single exercise, Reason: %v\n", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
	return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Exercise deleted successfully",
	})
}

func DeleteAllExercise(c *gin.Context) {
	
	result := gormDBConnect.Where("id LIKE ?", " ").Delete(Exercise{})

	if result.Error != nil {
		log.Printf("Error while deleting all exercise entries, Reason: %v\n", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
	return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Exercises deleted successfully",
	})
}