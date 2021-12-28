package routes

import (
	"os"

	"github.com/EfrenAL/simplexfit/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/", controllers.MainPage)

	//Exercises
	router.GET("/exercise",  JWTAuthMiddleware(), controllers.GetAllExercise)
	router.GET("/exercise/:exerciseId", JWTAuthMiddleware(),controllers.GetSingleExercise)
	router.POST("/exercise", JWTAuthMiddleware(), controllers.CreateExercise)
	router.POST("/exercises", JWTAuthMiddleware(), controllers.CreateExerciseBatch)
	router.DELETE("/exercise/:exerciseId", JWTAuthMiddleware(), controllers.DeleteExercise)	
	//Workouts
	router.GET("/workout", JWTAuthMiddleware(), controllers.GetAllWorkout)
	router.GET("/workout/:workoutId", JWTAuthMiddleware(), controllers.GetSingleWorkout)
	router.POST("/workout", JWTAuthMiddleware(), controllers.CreateWorkout)
	router.DELETE("/workout/:workoutId", JWTAuthMiddleware(), controllers.DeleteWorkout)	
	//User
	router.GET("/user", JWTAuthMiddleware(), controllers.GetUser)

}

//JWTAuthMiddleware middleware
func JWTAuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.Request.Header.Get("API-KEY")

		if checkToken(token) {
			c.Next()
		} else {
			c.AbortWithStatus(401)
		}
		c.Next()
	}
}

func checkToken(token string) bool {

	apiKey := os.Getenv("APIKEY")
    if token == apiKey {
		return true
	} else {
		return false
	}		
}

