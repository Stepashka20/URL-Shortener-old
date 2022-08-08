package main

import (
	"github.com/joho/godotenv"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/gin-gonic/contrib/static"
	"urlShorter/routes"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := routes.Init()
	router.Use(static.Serve("/", static.LocalFile("./static", true)))
	router.UrlRoute()
	router.Run("localhost:" + os.Getenv("PORT"))
}
