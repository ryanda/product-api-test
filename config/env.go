package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var ServicePort = "3000"
var BackendUrl string
var ConsumerKey string
var ConsumerSecret string

func loadEnv() {
	err = godotenv.Load()
	if err != nil {
		fmt.Printf("init.Env: Cannot load .env (%s)\n", err)
	}

	// backend url
	BackendUrl = os.Getenv("BACKEND_URL")

	// service config
	ServicePort = os.Getenv("APP_PORT")

	// creds
	ConsumerKey = os.Getenv("CONSUMER_KEY")
	ConsumerSecret = os.Getenv("CONSUMER_SECRET")
}
