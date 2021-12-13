package controllers

import (
	"database/sql"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

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

	result := gormDBConnect.Create(&Exercise{
		ID: string(rand.Intn(100)),
		Name: "Burpees",
		Repetitions: 5,
		Duration: 35,
		Complexity: "Hard",
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
	return


	/*var exercise Exercise
	c.BindJSON(&exercise)

	name := exercise.Name
	repetitions := exercise.Repetitions
	duration := exercise.Duration
	complexity := exercise.Complexity
	id := guuid.New().String()

	insertError := dbConnect.Insert(&Exercise{
		ID: id,
		Name: name,
		Repetitions: repetitions,
		Duration: duration,
		Complexity: complexity,
		CreatedAt: time.Now(),
		UpdatedAt:time.Now(),
	})
	if insertError != nil {
		log.Printf("Error while inserting new todo into db, Reason: %v\n", insertError)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Todo created Successfully",
	})
	return*/
}

func GetSingleExercise(c *gin.Context) {
	/*exerciseId := c.Param("exerciseId")
	exercise := &Exercise{ID: exerciseId}
	err := dbConnect.Select(exercise)
	if err != nil {
		log.Printf("Error while getting a single todo, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Todo not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Single Todo",
		"data": exercise,
	})
	return*/
}

func DeleteExercise(c *gin.Context) {
	/*exerciseId := c.Param("exerciseId")
	exercise := &Exercise{ID: exerciseId}
	err := dbConnect.Delete(exercise)
	if err != nil {
		log.Printf("Error while deleting a single todo, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
	return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Todo deleted successfully",
	})
	return*/
}