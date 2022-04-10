package service

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/prononciation2/stores"
)

// Config ADD DB CONFIG HERE LATER
type Config struct {
	SvcHost string
	SvcPort int
}

func Start(cfg *Config) {
	router := gin.New()

	db := stores.NewMongoDB("prononciation")
	s := NewServer(router, db)
	if err := s.Router.Run(fmt.Sprintf("%s:%d", cfg.SvcHost, cfg.SvcPort)); err != nil {
		log.Fatal(context.Background(), err.Error())
	}
}
