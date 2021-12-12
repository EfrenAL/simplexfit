package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
	orm "github.com/go-pg/pg/v9/orm"
	guuid "github.com/google/uuid"
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

var dbConnect *pg.DB

// Create User Table
func CreateExerciseTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.CreateTable(&Exercise{}, opts)

	if createError != nil {
		log.Printf("Error while creating todo table, Reason: %v\n", createError)
		return createError
	}
	
	log.Printf("Todo table created")
	return nil
}


func InitiateDB(db *pg.DB) {
	dbConnect = db
}

func GetAllExercise(c *gin.Context) {
	var exercise []Exercise
	err := dbConnect.Model(&exercise).Select()
	if err != nil {
		log.Printf("Error while getting all exercises, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Todos",
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
	return
}

func GetSingleExercise(c *gin.Context) {
	exerciseId := c.Param("exerciseId")
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
	return
}

func DeleteExercise(c *gin.Context) {
	exerciseId := c.Param("exerciseId")
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
	return
}