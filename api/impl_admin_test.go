package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var TestClientIdList = []string{"2dd54ac2-9772-781f-aaf5-e9a7a7bea73d"}

func TestClientList(t *testing.T) {
	router := GetTestRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/admin/client", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotImplemented, w.Code)
}

func TestClientCreate(t *testing.T) {
	router := GetTestRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/admin/client", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotImplemented, w.Code)
}

func TestClientDelete(t *testing.T) {
	router := GetTestRouter()

	w := httptest.NewRecorder()
	for _, client_id := range TestClientIdList {
		req, _ := http.NewRequest("POST", "/admin/client/"+client_id, nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotImplemented, w.Code)
	}
}

func TestClientRevoke(t *testing.T) {
	router := GetTestRouter()

	w := httptest.NewRecorder()
	for _, client_id := range TestClientIdList {
		req, _ := http.NewRequest("POST", "/admin/client/"+client_id+"/revoke", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotImplemented, w.Code)
	}
}

func TestExtInfoCreate(t *testing.T) {
	router := GetTestRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/admin/extinfo", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotImplemented, w.Code)
}

func TestExtInfoDelete(t *testing.T) {
	router := GetTestRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/admin/extinfo/{extinfo_id}", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotImplemented, w.Code)
}

func TestExtInfoRevoke(t *testing.T) {
	router := GetTestRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/admin/extinfo/{extinfo_id}/revoke", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotImplemented, w.Code)
}

func TestKoyoCreate(t *testing.T) {
	router := GetTestRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/admin/koyo", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotImplemented, w.Code)
}

func TestKoyoDelete(t *testing.T) {
	router := GetTestRouter()

	w := httptest.NewRecorder()
	for _, koyo_id := range TestKoyoIdList {
		req, _ := http.NewRequest("DELETE", "/admin/koyo/"+koyo_id, nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotImplemented, w.Code)
	}
}

func TestKoyoRevoke(t *testing.T) {
	router := GetTestRouter()

	w := httptest.NewRecorder()
	for _, koyo_id := range TestKoyoIdList {
		req, _ := http.NewRequest("POST", "/admin/koyo/"+koyo_id+"/revoke", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotImplemented, w.Code)
	}
}
