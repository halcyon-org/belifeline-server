package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDataList(t *testing.T) {
	router := GetTestRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/provider/data", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotImplemented, w.Code)
}

func TestEachDataGet(t *testing.T) {
	router := GetTestRouter()

	w := httptest.NewRecorder()

	for _, koyo_id := range TestKoyoIdList {
		req, _ := http.NewRequest("GET", "/provider/data/"+koyo_id+"?type=GeoJSON", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNotImplemented, w.Code)
	}
}
