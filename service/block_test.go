package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/oschwald/geoip2-golang"
)

func TestBlock_Blocked(t *testing.T) {
	db, _ = geoip2.Open("GeoLite2-Country.mmdb")
	defer db.Close()
	req, err := http.NewRequest("POST", "/", strings.NewReader(`{"ip":"1.0.0.1", "countries":["US"]}`))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(block)

	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Handler return wrong status code: got %v instead of: %v", recorder.Code, http.StatusOK)
	}

	expected := `{"action":"BLOCK"}`
	if recorder.Body.String() != expected {
		t.Errorf("Hanlder returned unexpected body: got %v instead of: %v", recorder.Body.String(), expected)
	}

}

func TestBlock_Allowed(t *testing.T) {
	db, _ = geoip2.Open("GeoLite2-Country.mmdb")
	defer db.Close()
	req, err := http.NewRequest("POST", "/", strings.NewReader(`{"ip":"1.0.0.1", "countries":["US", "AU"]}`))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(block)

	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Handler return wrong status code: got %v instead of: %v", recorder.Code, http.StatusOK)
	}

	expected := `{"action":"ALLOW"}`
	if recorder.Body.String() != expected {
		t.Errorf("Hanlder returned unexpected body: got %v instead of: %v", recorder.Body.String(), expected)
	}
}

func TestBlock_BadRequest(t *testing.T) {
	db, _ = geoip2.Open("GeoLite2-Country.mmdb")
	defer db.Close()
	req, err := http.NewRequest("POST", "/", strings.NewReader(`{"ip":"1.0.0.1", "countries":"US"}`))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(block)

	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusBadRequest {
		t.Errorf("Handler return wrong status code: got %v instead of: %v", recorder.Code, http.StatusBadRequest)
	}
}
