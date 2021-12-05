package router

import (
	"github.com/gorilla/mux"
	"github.com/heroku/go-getting-started/middleware"
)

// Router is exported and used in main.go
func Router() *mux.Router {

    router := mux.NewRouter()
	router.HandleFunc("/api/user", middleware.GetAllExercises).Methods("GET", "OPTIONS")
    
    return router
}