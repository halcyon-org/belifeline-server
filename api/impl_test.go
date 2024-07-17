package api_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/halcyon-org/kizuna/api"
	"github.com/stretchr/testify/assert"
)

func GetTestRouter() *gin.Engine {
	gin.DefaultWriter = io.Discard

	server := api.NewServer()
	r := gin.Default()

	api.RegisterHandlers(r, server)

	return r
}

func TestStatusGet(t *testing.T) {
	router := GetTestRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/status", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
