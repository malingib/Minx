package config

import "testing"

func TestLoadDefaults(t *testing.T) {
	t.Setenv("MINX_SERVICE_NAME", "")
	t.Setenv("MINX_HTTP_ADDR", "")
	got, err := Load()
	if err != nil { t.Fatalf("Load() error = %v", err) }
	if got.ServiceName != "minx-core-api" { t.Fatalf("ServiceName = %q", got.ServiceName) }
	if got.HTTPAddr != ":8080" { t.Fatalf("HTTPAddr = %q", got.HTTPAddr) }
}

func TestLoadRejectsInvalidAddress(t *testing.T) {
	t.Setenv("MINX_HTTP_ADDR", "not-an-address")
	if _, err := Load(); err == nil { t.Fatal("expected invalid address error") }
}
