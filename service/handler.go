package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prononciation2/models"
	"github.com/prononciation2/stores"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Server struct {
	CityService Service
	Router      *gin.Engine
	Logger      *zap.SugaredLogger
}

// NewServer Here I should add a db as param
func NewServer(router *gin.Engine, db *mongo.Database, logger *zap.SugaredLogger) Server {
	cityService := NewService(stores.NewMongoCityRepo(db))
	s := Server{
		CityService: cityService,
		Router:      router,
		Logger:      logger,
	}
	s.Router.GET("/cities/", s.GetCities)
	s.Router.GET("/cities/:id", s.GetCityByID)
	s.Router.POST("/cities/", s.CreateCity)
	s.Router.DELETE("/cities/:id", s.DeleteCity)
	s.Router.PUT("/cities/:id", s.UpdateCity)
	s.Router.GET("/test", s.Test)

	return s
}

// ADD ALL THE HANDLERS HERE

// GetCities returns a list of all cities
func (s *Server) GetCities(c *gin.Context) {
	cities, err := s.CityService.GetCities(c)
	if err != nil {
		s.Logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cities)
}

// GetCityByID returns a specific city
func (s *Server) GetCityByID(c *gin.Context) {
	city, err := s.CityService.GetCityByID(c, c.Param("id"))
	if err != nil {
		s.Logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, city)
}

// CreateCity creates a city
func (s *Server) CreateCity(c *gin.Context) {
	s.Logger.Info(c.Request.URL.Path, zap.String("method", c.Request.Method))
	var city models.City
	if err := c.ShouldBindJSON(&city); err != nil {
		s.Logger.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	city, err := s.CityService.CreateCity(c, city)
	if err != nil {
		s.Logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, city)
}

// DeleteCity deletes a specific city
func (s *Server) DeleteCity(c *gin.Context) {
	err := s.CityService.DeleteCity(c, c.Param("id"))
	if err != nil {
		s.Logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusAccepted)
}

// UpdateCity updates a specific city
func (s *Server) UpdateCity(c *gin.Context) {
	var city models.City
	if err := c.ShouldBindJSON(&city); err != nil {
		s.Logger.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	city, err := s.CityService.UpdateCity(c, city)
	if err != nil {
		s.Logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, city)
}

// Test DELME
func (s *Server) Test(c *gin.Context) {
	c.JSON(http.StatusOK, "Working")
}
