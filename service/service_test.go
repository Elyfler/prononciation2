package service

import (
	"context"
	"testing"

	"github.com/prononciation2/models"
	"github.com/prononciation2/stores"
)

func TestCreateCity(t *testing.T) {

	db, _ := stores.NewMongoDB("test")
	repo := stores.NewMongoCityRepo(db)
	s := NewService(repo)

	cityInput := models.City{
		Name:          "TestName",
		PostCode:      "29200",
		Prononciation: "TestName",
	}
	cityOutput, _ := s.CreateCity(context.Background(), cityInput)

	if cityOutput.Name != cityInput.Name {
		t.Errorf("Expected: %v, got: %v", cityInput.Name, cityOutput.Name)
	}

	if cityOutput.PostCode != cityInput.PostCode {
		t.Errorf("Expected: %v, got: %v", cityInput.PostCode, cityOutput.PostCode)
	}

	if cityOutput.Prononciation != cityInput.Prononciation {
		t.Errorf("Expected: %v, got: %v", cityInput.Prononciation, cityOutput.Prononciation)
	}

	if cityOutput.ID == "" {
		t.Error("Expected city to have an ID")
	}
}
