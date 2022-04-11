package service

import (
	"context"

	"github.com/prononciation2/models"
	"github.com/prononciation2/stores"
)

type Service interface {
	CreateCity(ctx context.Context, c models.City) (models.City, error)
	GetCityByID(ctx context.Context, id string) (models.City, error)
	GetCities(ctx context.Context) ([]models.City, error)
	DeleteCity(ctx context.Context, id string) error
	UpdateCity(ctx context.Context, c models.City) (models.City, error)
}

type CityService struct {
	repository stores.CityRepo
}

func NewService(repo stores.CityRepo) Service {
	return &CityService{repo}
}

func (s CityService) CreateCity(ctx context.Context, c models.City) (models.City, error) {
	return s.repository.CreateCity(ctx, c)
}
func (s CityService) GetCityByID(ctx context.Context, id string) (models.City, error) {
	return s.repository.GetCityByID(ctx, id)
}
func (s CityService) GetCities(ctx context.Context) ([]models.City, error) {
	return s.repository.GetCities(ctx)
}
func (s CityService) DeleteCity(ctx context.Context, id string) error {
	return s.repository.DeleteCity(ctx, id)
}
func (s CityService) UpdateCity(ctx context.Context, c models.City) (models.City, error) {
	return s.repository.UpdateCity(ctx, c)
}
