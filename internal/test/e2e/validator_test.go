package e2e

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net/http"
	"testing"

	"github.com/halcyon-org/kizuna/internal/adapter/interceptor"
	"github.com/halcyon-org/kizuna/internal/test/helper"
	"github.com/stretchr/testify/assert"
)

func TestValidClientSet(t *testing.T) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	apiKey := helper.GenAPIKey(t)

	s := helper.NewServer(t)

	data := "{\"username\":\"test\"}"

	url := fmt.Sprintf("%s/belifeline.v1.AdminService/ClientSet", s.URL)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(data)))
	if err != nil {
		t.Fatalf("http.NewRequest() = %v", err)
	}
	req.Header.Add(interceptor.AuthAPIKeyHeader, apiKey)
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("http.DefaultClient.Do() = %v", err)
	}
	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestInvalidClientSet(t *testing.T) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	apiKey := helper.GenAPIKey(t)

	s := helper.NewServer(t)

	data := "{\"name\":\"test\"}"

	url := fmt.Sprintf("%s/belifeline.v1.AdminService/ClientSet", s.URL)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(data)))
	if err != nil {
		t.Fatalf("http.NewRequest() = %v", err)
	}
	req.Header.Add(interceptor.AuthAPIKeyHeader, apiKey)
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("http.DefaultClient.Do() = %v", err)
	}
	defer res.Body.Close()

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}
