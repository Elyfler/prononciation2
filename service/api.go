package service

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/prononciation2/logger"
	"github.com/prononciation2/stores"
)

// Config ADD DB CONFIG HERE LATER
type Config struct {
	SvcHost string
	SvcPort int
}

func Start(cfg *Config) {
	router := gin.New()

	db, err := stores.NewMongoDB("prononciation")
	if err != nil {
		log.Fatal(err)
	}
	logger := logger.NewLogger()
	s := NewServer(router, db, logger)
	if err := s.Router.Run(fmt.Sprintf("%s:%d", cfg.SvcHost, cfg.SvcPort)); err != nil {
		log.Fatal(context.Background(), err.Error())
	}
}
