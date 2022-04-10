package service

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetCities(t *testing.T) {
	router := gin.Default()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	s := NewServer(router)

	s.Router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Error("Did not work")
	}

}
