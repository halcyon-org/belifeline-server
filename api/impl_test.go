package api

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func getRouter() *gin.Engine {
	gin.DefaultWriter = io.Discard

	server := NewServer()
	r := gin.Default()

	RegisterHandlers(r, server)

	return r
}

func TestStatusRoute(t *testing.T) {
	router := getRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/status", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
