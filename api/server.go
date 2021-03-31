package api

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/kartuludus/wallaster-be/api/controllers"
	"github.com/kartuludus/wallaster-be/api/seed"
)

var server = controllers.Server{}

func Run() {

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	server.Initialize(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	shouldReseed, err := strconv.ParseBool(os.Getenv("RESEED"))
	if (shouldReseed) {
		fmt.Println("We are reseeding the db")
		seed.Load(server.DB)
	}

	server.Run(":8080")

}