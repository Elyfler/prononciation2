package stores

import (
	"context"

	"github.com/prononciation2/models"
)

type CityRepo interface {
	CreateCity(ctx context.Context, c models.City) (models.City, error)
	GetCityByID(ctx context.Context, id string) (models.City, error)
	GetCities(ctx context.Context) ([]models.City, error)
	DeleteCity(ctx context.Context, id string) error
	UpdateCity(ctx context.Context, c models.City) (models.City, error)
}
