package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/EfrenAL/SimpleXfit/controllers"
)
func Routes(router *gin.Engine) {
	router.GET("/", controllers.MainPage)
	router.GET("/exercise", controllers.GetAllExercise)
	router.GET("/exercise/:exerciseId", controllers.GetSingleExercise)

	router.POST("/exercise", controllers.CreateExercise)
	router.POST("/form/exercise", controllers.CreateExerciseForm)
	router.POST("/exercises", controllers.CreateExerciseBatch)

	router.DELETE("/exercise/:exerciseId", controllers.DeleteExercise)
	router.DELETE("/exercise/all", controllers.DeleteExercise)
}