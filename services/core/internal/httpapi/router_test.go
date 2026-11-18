package httpapi

import (
	"errors"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func testLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(io.Discard, nil))
}

func TestHealthEndpoints(t *testing.T) {
	handler := NewRouter(testLogger(), "test-core", func(_ context.Context) error { return nil })
	for _, path := range []string{"/health/live", "/health/ready"} {
		t.Run(path, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, path, nil)
			handler.ServeHTTP(rec, req)

			if rec.Code != http.StatusOK {
				t.Fatalf("status = %d", rec.Code)
			}
			if !strings.Contains(rec.Body.String(), `"service":"test-core"`) {
				t.Fatalf("unexpected body: %s", rec.Body.String())
			}
			if rec.Header().Get("X-Content-Type-Options") != "nosniff" {
				t.Fatal("missing security header")
			}
		})
	}
}

func TestReadinessFailsWhenDependencyIsUnavailable(t *testing.T) {
	handler := NewRouter(testLogger(), "test-core", func(_ context.Context) error {
		return errors.New("database unavailable")
	})
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/health/ready", nil))

	if rec.Code != http.StatusServiceUnavailable {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusServiceUnavailable)
	}
	if !strings.Contains(rec.Body.String(), `"status":"not_ready"`) {
		t.Fatalf("unexpected body: %s", rec.Body.String())
	}
}

func TestLivenessDoesNotDependOnDatabaseReadiness(t *testing.T) {
	handler := NewRouter(testLogger(), "test-core", func(_ context.Context) error {
		return errors.New("database unavailable")
	})
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/health/live", nil))

	if rec.Code != http.StatusOK {
		t.Fatalf("liveness status = %d, want %d", rec.Code, http.StatusOK)
	}
}

func TestReadinessWithoutProbeIsUnavailable(t *testing.T) {
	handler := NewRouter(testLogger(), "test-core", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/health/ready", nil))
	if rec.Code != http.StatusServiceUnavailable {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusServiceUnavailable)
	}
}

func TestUnknownRouteIsNotFound(t *testing.T) {
	handler := NewRouter(testLogger(), "test-core", func(_ context.Context) error { return nil })
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/missing", nil))

	if rec.Code != http.StatusNotFound {
		t.Fatalf("status = %d", rec.Code)
	}
}
