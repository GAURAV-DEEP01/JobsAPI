package main

import (
	"github.com/gaurav-deep01/jobboard-api/internal/db"
	"github.com/gaurav-deep01/jobboard-api/internal/routes"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Failed to load .env file")
	}
	db.InitMongo(os.Getenv("DATABASE_URL"))
	r := routes.SetupRouter()

	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, r)
}
