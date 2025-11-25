package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"Template_Go/internal/controllers"
)

func TestHomeHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	controllers.HomeHandler(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("attendu %d, obtenu %d", http.StatusOK, resp.StatusCode)
	}
}