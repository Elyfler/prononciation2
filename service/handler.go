package service

import (
	"github.com/gin-gonic/gin"
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

func (h *handler) GetCities(c *gin.Context)   {}
func (h *handler) GetCityByID(c *gin.Context) {}
func (h *handler) CreateCity(c *gin.Context)  {}
func (h *handler) DeleteCity(c *gin.Context)  {}
func (h *handler) UpdateCity(c *gin.Context)  {}
