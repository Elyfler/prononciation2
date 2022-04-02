package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/prononciation2/service"
)

func main() {
	fmt.Println("Hello, world !")

	port, err := strconv.Atoi(os.Getenv("CITIES_PORT"))
	if err != nil {
		log.Fatal("Could not retrieve port")
	}
	cfg := service.Config{
		SvcHost: os.Getenv("CITIES_HOST"),
		SvcPort: port,
	}
	service.Start(&cfg)
}
