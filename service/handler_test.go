package service

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/prononciation2/logger"
	"github.com/prononciation2/stores"
)

func setupTestServer() Server {
	router := gin.Default()
	db, _ := stores.NewMongoDB("test")
	logger := logger.NewLogger()
	return NewServer(router, db, logger)
}

func TestTest(t *testing.T) {

	s := setupTestServer()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	s.Router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Error("Did not work")
	}
}
