# simplexfit

Xfit

## Running Locally

You need to install postgresinyour local machine and use your user and password
Use this configuration to run it in production:
//Local env
conection := "user=YOUR-USER-HERE dbname=postgres password=YOUR-PASSWORD-HERE host=localhost sslmode=disable"
port := "8080"

## Deploying to Heroku

```sh
$ heroku create
$ git push heroku main
$ heroku open
```

Use this configuration to run it in production:
//Production env
conection := os.Getenv("DATABASE_URL")
port := os.Getenv("PORT")

## Documentation

More info to follow
