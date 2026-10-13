package httpapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthz(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rec := httptest.NewRecorder()

	NewRouter().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("Код ответа = %d. Был ожидаем код ответа: %d", rec.Code, http.StatusOK)
	}

	if ct := rec.Header().Get("Content-Type"); ct != "application/json" {
		t.Errorf("Content-Type = %q. Был ожидаем Content-Type=application/json", ct)
	}

	var body healthResponse

	if err := json.NewDecoder(rec.Body).Decode(&body); err != nil {
		t.Fatalf("Тело ответа не формата JSON: %v", err)
	}

	if body.Status != "ok" {
		t.Errorf("status = %q. Был ожидаем status=ok", body.Status)
	}
}

func TestHealthzWrongMethod(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/healthz", nil)
	rec := httptest.NewRecorder()

	NewRouter().ServeHTTP(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("Код ответа = %d. Был ожидаем код ответа: %d", rec.Code, http.StatusMethodNotAllowed)
	}
}
