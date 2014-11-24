package main_test

import (
  "net/http"
  "net/http/httptest"
  "testing"
)

var (
  server *httptest.Server
  reader io.Reader
  purchaseUrl string
)

func init() {
  server = httptest.NewServer(
}
