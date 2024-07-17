package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var TestKoyoIdList = []string{"123", "456", "789"}

func TestKoyoList(t *testing.T) {
	router := GetTestRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/koyo", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotImplemented, w.Code)
}

func TestEachKoyoGet(t *testing.T) {
	router := GetTestRouter()

	w := httptest.NewRecorder()
	for _, koyo_id := range TestKoyoIdList {
		req, _ := http.NewRequest("GET", "/koyo/"+koyo_id, nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotImplemented, w.Code)
	}
}

func TestEachKoyoUpdate(t *testing.T) {
	router := GetTestRouter()

	w := httptest.NewRecorder()
	for _, koyo_id := range TestKoyoIdList {
		req, _ := http.NewRequest("PUT", "/koyo/"+koyo_id, nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotImplemented, w.Code)
	}
}

func TestEachKoyoDataNew(t *testing.T) {
	router := GetTestRouter()

	w := httptest.NewRecorder()
	for _, koyo_id := range TestKoyoIdList {
		req, _ := http.NewRequest("POST", "/koyo/"+koyo_id+"/data", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotImplemented, w.Code)
	}
}
