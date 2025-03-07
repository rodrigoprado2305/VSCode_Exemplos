package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTesteHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(teste)

	handler.ServeHTTP(rr, req)

	expected := "<h1>Teste 1070</h1>"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
