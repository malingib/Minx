package httpapi

import (
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealthEndpoints(t *testing.T) {
	handler := NewRouter(slog.New(slog.NewTextHandler(io.Discard, nil)), "test-core")
	for _, path := range []string{"/health/live", "/health/ready"} {
		t.Run(path, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, path, nil)
			handler.ServeHTTP(rec, req)
			if rec.Code != http.StatusOK { t.Fatalf("status = %d", rec.Code) }
			if !strings.Contains(rec.Body.String(), `"service":"test-core"`) { t.Fatalf("unexpected body: %s", rec.Body.String()) }
			if rec.Header().Get("X-Content-Type-Options") != "nosniff" { t.Fatal("missing security header") }
		})
	}
}

func TestUnknownRouteIsNotFound(t *testing.T) {
	handler := NewRouter(slog.New(slog.NewTextHandler(io.Discard, nil)), "test-core")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/missing", nil))
	if rec.Code != http.StatusNotFound { t.Fatalf("status = %d", rec.Code) }
}
