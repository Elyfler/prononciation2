package service

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/prononciation2/stores"
)

func TestGetCities(t *testing.T) {
	router := gin.Default()
	db, _ := stores.NewMongoDB("test")
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	s := NewServer(router, db)

	s.Router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Error("Did not work")
	}

}
