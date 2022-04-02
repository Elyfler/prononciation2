package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prononciation2/models"
	"github.com/prononciation2/stores"
)

type handler struct {
	CityService Service
}

// Activate Here I should add a db as param
func Activate(router *gin.Engine) {
	cityService := NewService(stores.NewMongoCityRepo())
	newHandler(router, cityService)
}

func newHandler(router *gin.Engine, cs Service) {
	h := handler{
		CityService: cs,
	}
	router.GET("/cities/", h.GetCities)
	router.GET("/cities/:id", h.GetCityByID)
	router.POST("/cities/", h.CreateCity)
	router.DELETE("/cities/:id", h.DeleteCity)
	router.PUT("/cities/:id", h.UpdateCity)
}

// ADD ALL THE HANDLERS HERE

// GetCities returns a list of all cities
func (h *handler) GetCities(c *gin.Context) {
	cities, err := h.CityService.GetCities(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cities)
}

// GetCityByID returns a specific city
func (h *handler) GetCityByID(c *gin.Context) {
	city, err := h.CityService.GetCityByID(c, c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, city)
}

// CreateCity creates a city
func (h *handler) CreateCity(c *gin.Context) {
	var city models.City
	if err := c.ShouldBindJSON(&city); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	city, err := h.CityService.CreateCity(c, city)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, city)
}

// DeleteCity deletes a specific city
func (h *handler) DeleteCity(c *gin.Context) {
	err := h.CityService.DeleteCity(c, c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusAccepted)
}

// UpdateCity updates a specific city
func (h *handler) UpdateCity(c *gin.Context) {
	var city models.City
	if err := c.ShouldBindJSON(&city); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	city, err := h.CityService.UpdateCity(c, city)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, city)
}
