package main

import (
	"log"

	config "github.com/EfrenAL/simplexfit/configs"
	routes "github.com/EfrenAL/simplexfit/router"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	_ "github.com/lib/pq"
)

func main() {
    
	
	config.Connect()
	
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.tmpl.html")
    router.Static("/static", "static")

	routes.Routes(router)
	log.Fatal(router.Run())
}