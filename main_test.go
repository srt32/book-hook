package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	server      *httptest.Server
	reader      io.Reader
	purchaseUrl string
)

func init() {
	server = httptest.NewServer(Handlers())

	purchaseUrl = fmt.Sprintf("%s/purchase", server.URL)
}

func TestCreatePurchase(t *testing.T) {
	purchaseJson := `{"email": "me@example.com", "product_id": 100}`

	reader = strings.NewReader(purchaseJson)

	request, err := http.NewRequest("POST", purchaseUrl, reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("Success expected, got: %d", res.StatusCode)
	}
}
