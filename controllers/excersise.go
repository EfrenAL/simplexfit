package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//"gorm.io/gorm"


type Exercise struct {
	gorm.Model
	Name 			string `json:"name"`
	Repetitions 	int64  `json:"repetitions"`
	Weight 			int64  `json:"weight"`
	Complexity 		string `json:"complexity"`
}

// Create User Table
func PopulateExerciseTable() error {

	//gormDBConnect.Migrator().CreateTable(&Exercise{})

	Custs1 := Exercise{Name: "John",Repetitions: 10, Weight: 20, Complexity: "hard" } 
	Custs2 := Exercise{Name: "John2",Repetitions: 10, Weight: 20, Complexity: "hard2" } 

	gormDBConnect.Create(&Custs1)
	gormDBConnect.Create(&Custs2)



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
	weight := exercise.Weight
	complexity := exercise.Complexity

	result := gormDBConnect.Create(&Exercise{	
		Name: name,
		Repetitions: repetitions,
		Weight: weight,
		Complexity: complexity,
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

	result := gormDBConnect.Create(&Exercise{
		Name: name,
		Repetitions: 100,
		Weight: 40,
		Complexity: complexity,
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
	exercise := &Exercise{}
	result := gormDBConnect.First(&exercise, exerciseId)
	
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
	exercise := &Exercise{}

	result := gormDBConnect.Delete(exercise, exerciseId)

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