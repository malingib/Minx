package config

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	ServiceName string
	HTTPAddr    string
}

func Load() (Config, error) {
	cfg := Config{
		ServiceName: strings.TrimSpace(os.Getenv("MINX_SERVICE_NAME")),
		HTTPAddr:    strings.TrimSpace(os.Getenv("MINX_HTTP_ADDR")),
	}
	if cfg.ServiceName == "" {
		cfg.ServiceName = "minx-core-api"
	}
	if cfg.HTTPAddr == "" {
		cfg.HTTPAddr = ":8080"
	}
	if !strings.HasPrefix(cfg.HTTPAddr, ":") && !strings.Contains(cfg.HTTPAddr, ":") {
		return Config{}, fmt.Errorf("MINX_HTTP_ADDR must be a host:port address")
	}
	return cfg, nil
}
