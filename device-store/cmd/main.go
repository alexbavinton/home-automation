package main

import (
	"os"

	"github.com/alexbavinton/home-automation/device-store/internal/service"
)

func main() {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	service.Run(redisHost, redisPort)
}
