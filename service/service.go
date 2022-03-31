package service

import (
	"context"

	"github.com/prononciation2/models"
	"github.com/prononciation2/stores"
)

// Add logger here
type Service struct {
	repository stores.CityRepo
}

func NewService(repo stores.CityRepo) *Service {
	return &Service{
		repository: repo,
	}
}

func (s Service) CreateCity(ctx context.Context, c models.City) (models.City, error) {
	return s.repository.CreateCity(ctx, c)
}
func (s Service) GetCityByID(ctx context.Context, id string) (models.City, error) {
	return s.repository.GetCityByID(ctx, id)
}
func (s Service) GetCities(ctx context.Context) ([]models.City, error) {
	return s.repository.GetCities(ctx)
}
func (s Service) DeleteCity(ctx context.Context, id string) error {
	return s.repository.DeleteCity(ctx, id)
}
func (s Service) UpdateCity(ctx context.Context, c models.City) (models.City, error) {
	return s.repository.UpdateCity(ctx, c)
}
