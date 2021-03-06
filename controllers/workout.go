package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

type Workout struct {
	gorm.Model
	Name string    `json:"name"`
	Date time.Time `json:"date"`
	//Excercises 		Exercise  `json:"excercises" gorm:"ForeignKey:ExerciseID;references:ID"`
	Exercises  []Exercise `json:"excercises" gorm:"many2many:workout_exercise;"`
	ExerciseID uint       `json:"-"`
}

// Create User Table
func PopulateWorkoutTable() error {

	Custs1 := Workout{Name: "WOD Xfit Amsterdam", Exercises: []Exercise{
		{Name: "Flexiones", Repetitions: 10, Weight: 20, Complexity: "hard5"},
		{Name: "Abs", Repetitions: 10, Weight: 20, Complexity: "hard5"}},
	}
	gormDBConnect.Create(&Custs1)

	log.Printf("Workout table populated")
	return nil
}

func GetAllWorkout(c *gin.Context) {
	var workout []Workout

	result := gormDBConnect.Preload("Exercises").Find(&workout)

	if result.Error != nil {
		log.Printf("Error while getting all workout, Reason: %v\n", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}
	message := "All Exercise %s" + string(result.RowsAffected)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": message,
		"data":    workout,
	})
}

func GetSingleWorkout(c *gin.Context) {
	workoutId := c.Param("workoutId")
	workout := &Workout{}

	result := gormDBConnect.First(&workout, workoutId)

	if result.Error != nil {
		log.Printf("Error while getting a single workout, Reason: %v\n", result.Error)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Exercise not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Single Workout",
		"data":    workout,
	})
}

func CreateWorkout(c *gin.Context) {

	var workout Workout

	c.BindJSON(&workout)

	name := workout.Name
	date := workout.Date
	excercises := workout.Exercises

	result := gormDBConnect.Create(&Workout{
		Name:      name,
		Date:      date,
		Exercises: excercises,
	})

	if result.Error != nil {
		log.Printf("Error while inserting new workout into db, Reason: %v\n", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Workout created Successfully",
	})
}

func DeleteWorkout(c *gin.Context) {
	workoutId := c.Param("workoutId")
	workout := &Workout{}

	result := gormDBConnect.Delete(workout, workoutId)

	if result.Error != nil {
		log.Printf("Error while deleting a single workout, Reason: %v\n", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Workout deleted successfully",
	})
}
