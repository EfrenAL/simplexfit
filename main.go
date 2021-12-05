package main

import (
	"github.com/heroku/go-getting-started/router"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	_ "github.com/lib/pq"
)




func setupDB(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        if _, err := db.Exec("%s", queryCreateTable); err != nil {
            c.String(http.StatusInternalServerError,
                fmt.Sprintf("Error creating database table: %q", err))
            return
        }

        if _, err := db.Exec("INSERT INTO ticks VALUES (now())"); err != nil {
            c.String(http.StatusInternalServerError,
                fmt.Sprintf("Error incrementing tick: %q", err))
            return
        }

        rows, err := db.Query("SELECT tick FROM ticks")
        if err != nil {
            c.String(http.StatusInternalServerError,
                fmt.Sprintf("Error reading ticks: %q", err))
            return
        }

        defer rows.Close()
        for rows.Next() {
            var tick time.Time
            if err := rows.Scan(&tick); err != nil {
                c.String(http.StatusInternalServerError,
                    fmt.Sprintf("Error scanning ticks: %q", err))
                return
            }
            c.String(http.StatusOK, fmt.Sprintf("Read from DB: %s\n", tick.String()))
        }
    }
}


func main() {
    
	r := router.Router()
    fmt.Println("Starting server on the port 8080...")
    log.Fatal(http.ListenAndServe(":8080", r))

}