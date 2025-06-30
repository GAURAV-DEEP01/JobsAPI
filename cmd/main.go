package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gaurav-deep01/jobboard-api/internal/db"
	"github.com/gaurav-deep01/jobboard-api/internal/routes"
	"github.com/gaurav-deep01/jobboard-api/internal/util"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Failed to load .env file")
	}

	db.InitMongo(util.MustGetenv("DATABASE_URL"))
	r := routes.SetupRouter()

	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, r)
}
