package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"

	_ "github.com/heroku/x/hmetrics/onload"
	_ "github.com/lib/pq"
)

func main() {
    
	//r := router.Router()
    //fmt.Println("Starting server on the port 8080...")
    //log.Fatal(http.ListenAndServe(":8080", r))

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")


}



/*

func main() {
    port := os.Getenv("PORT")

    if port == "" {
        log.Fatal("$PORT must be set")
    }

    tStr := os.Getenv("REPEAT")
    repeat, err := strconv.Atoi(tStr)
    if err != nil {
        log.Printf("Error converting $REPEAT to an int: %q - Using default\n", err)
        repeat = 5
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

    router.GET("/mark", func(c *gin.Context) {
        //c.String(http.StatusOK, string(blackfriday.Run([]byte("**hi!**"))))
    })

    router.GET("/repeat", repeatHandler(repeat))

    router.GET("/db", dbFunc(db))

    router.Run(":" + port)
}

*/