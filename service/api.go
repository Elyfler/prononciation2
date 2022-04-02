package service

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

// Config ADD DB CONFIG HERE LATER
type Config struct {
	SvcHost string
	SvcPort int
}

func Start(cfg *Config) {
	router := gin.New()

	Activate(router)
	if err := router.Run(fmt.Sprintf("%s:%d", cfg.SvcHost, cfg.SvcPort)); err != nil {
		log.Fatal(context.Background(), err.Error())
	}
}