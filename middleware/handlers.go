package middleware

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/heroku/go-getting-started/models"
)

const (
	queryCreateTable = "CREATE TABLE exercise (exerciseId SERIAL PRIMARY KEY, name TEXT, repetitions INT, time INT, complexity TEXT);"
)

// create connection with postgres db
func createConnection() *sql.DB {
    port := os.Getenv("PORT")


    if port == "" {
        log.Fatal("$PORT must be set")
    }

    db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
    if err != nil {
        log.Fatalf("Error opening database: %q", err)
    }

	if _, err := db.Exec("%s", queryCreateTable); err != nil {
		log.Fatalf("Error creating database table: %q", err)
	}

	if _, err := db.Exec(`insert into "exercise"("exerciseId", "name", "repetitions","time", "complexity") values(1, 'Burpees', 5, 2, "high" )`); err != nil {		
		log.Fatalf("Error incrementing tick: %q", err)	
	}

    return db
}


// GetAllUser will return all the users
func GetAllExercises(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    // get all the users in the db
    users, err := getAllExercises()

    if err != nil {
        log.Fatalf("Unable to get all user. %v", err)
    }

    // send all the users as response
    json.NewEncoder(w).Encode(users)
}

// get one user from the DB by its userid
func getAllExercises() ([]models.Exercise, error) {
    // create the postgres db connection
    db := createConnection()

    // close the db connection
    defer db.Close()

    var exercises []models.Exercise

    // create the select sql query
    sqlStatement := `SELECT * FROM exercise`

    // execute the sql statement
    rows, err := db.Query(sqlStatement)

    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }

    // close the statement
    defer rows.Close()

    // iterate over the rows
    for rows.Next() {
        var exercise models.Exercise

        // unmarshal the row object to user
        err = rows.Scan(&exercise.ExerciseId, &exercise.Name, &exercise.Complexity, &exercise.Repetitions, &exercise.Time)

        if err != nil {
            log.Fatalf("Unable to scan the row. %v", err)
        }

        // append the user in the users slice
        exercises = append(exercises, exercise)

    }

    // return empty user on error
    return exercises, err
}

