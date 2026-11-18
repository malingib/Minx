package config

import "testing"

func TestLoadDefaults(t *testing.T) {
	t.Setenv("MINX_SERVICE_NAME", "")
	t.Setenv("MINX_HTTP_ADDR", "")
	t.Setenv("MINX_DATABASE_URL", "")

	got, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}
	if got.ServiceName != "minx-core-api" {
		t.Fatalf("ServiceName = %q", got.ServiceName)
	}
	if got.HTTPAddr != ":8080" {
		t.Fatalf("HTTPAddr = %q", got.HTTPAddr)
	}
	if got.DatabaseURL != "" {
		t.Fatalf("DatabaseURL = %q, want empty", got.DatabaseURL)
	}
}

func TestLoadDatabaseURL(t *testing.T) {
	t.Setenv("MINX_DATABASE_URL", "postgres://minx:secret@localhost:5432/minx")
	got, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}
	if got.DatabaseURL != "postgres://minx:secret@localhost:5432/minx" {
		t.Fatalf("DatabaseURL = %q", got.DatabaseURL)
	}
}

func TestLoadRejectsInvalidAddresses(t *testing.T) {
	for _, addr := range []string{
		"not-an-address",
		"localhost",
		"localhost:",
		"localhost:abc",
		"localhost:70000",
		":0",
	} {
		t.Run(addr, func(t *testing.T) {
			t.Setenv("MINX_HTTP_ADDR", addr)
			if _, err := Load(); err == nil {
				t.Fatalf("Load() accepted invalid address %q", addr)
			}
		})
	}
}

func TestLoadAcceptsValidAddresses(t *testing.T) {
	for _, addr := range []string{":8080", "127.0.0.1:8080", "localhost:9000", "[::1]:8080"} {
		t.Run(addr, func(t *testing.T) {
			t.Setenv("MINX_HTTP_ADDR", addr)
			if _, err := Load(); err != nil {
				t.Fatalf("Load() rejected %q: %v", addr, err)
			}
		})
	}
}
