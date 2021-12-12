package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/heroku/go-getting-started/controllers"
)
func Routes(router *gin.Engine) {
	router.GET("/", controllers.MainPage)
	router.GET("/exercise", controllers.GetAllExercise)
	router.POST("/exercise", controllers.CreateExercise)
	router.GET("/exercise/:exerciseId", controllers.GetSingleExercise)
	router.DELETE("/exercise/:exerciseId", controllers.DeleteExercise)
}

/*

c.HTML(http.StatusOK, "index.tmpl.html", nil)

package router

import (
	"github.com/gorilla/mux"
	
)

// Router is exported and used in main.go
func Router() *mux.Router {

    router := mux.NewRouter()
	router.HandleFunc("/api/exercises", middleware.GetAllExercises).Methods("GET", "OPTIONS")
	router.HandleFunc("/", middleware.GetAllExercises).Methods("GET", "OPTIONS")

    
    return router
}*/