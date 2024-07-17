package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtInfoList(t *testing.T) {
	router := GetTestRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/extinfo", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotImplemented, w.Code)
}

func TestEachExtInfoGet(t *testing.T) {
	router := GetTestRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/extinfo/{extinfo_id}", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotImplemented, w.Code)
}

func TestExampleInfoGet(t *testing.T) {
	router := GetTestRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/extinfo/example_id/data", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotImplemented, w.Code)
}

func TestExampleInfoPost(t *testing.T) {
	router := GetTestRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/extinfo/example_id/data", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotImplemented, w.Code)
}
