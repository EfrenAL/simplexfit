package main

import (
	"log"

	config "github.com/heroku/go-getting-started/configs"
	routes "github.com/heroku/go-getting-started/router"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	_ "github.com/lib/pq"
)

func main() {
    
	// Connect to db
	config.Connect()
	// Init Router
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.tmpl.html")
    router.Static("/static", "static")

	// Route Handlers / Endpoints
	routes.Routes(router)
	log.Fatal(router.Run())
}

	/*port := os.Getenv("PORT")

    if port == "" {
        log.Fatal("$PORT must be set")
    }

    
    db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
    if err != nil {
        log.Fatalf("Error opening database: %q", err)
    }

    router := gin.New()
    router.Use(gin.Logger())
    router.LoadHTMLGlob("templates/*.tmpl.html")
    router.Static("/static", "static")

    router.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.tmpl.html", nil)
    })

    router.GET("/db", dbFunc(db))

    router.Run(":" + port) */