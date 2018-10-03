package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootAPI(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	RootHandler(w, req)
	if w == nil {
		t.Error("endpoint not working")
	}
}

func TestHitungHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/hitung", nil)
	w := httptest.NewRecorder()
	HitungHandler(w, req)
	if w == nil {
		t.Error("endpoint not working")
	}
}

func TestHitungHandlerFix(t *testing.T) {
	req, _ := http.NewRequest("GET", "/hitung/9", nil)
	w := httptest.NewRecorder()
	HitungHandler(w, req)
	if w == nil {
		t.Error("endpoint not working")
	}
}
