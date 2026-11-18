package config

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	ServiceName string
	HTTPAddr    string
	DatabaseURL string
}

func Load() (Config, error) {
	cfg := Config{
		ServiceName: strings.TrimSpace(os.Getenv("MINX_SERVICE_NAME")),
		HTTPAddr:    strings.TrimSpace(os.Getenv("MINX_HTTP_ADDR")),
		DatabaseURL: strings.TrimSpace(os.Getenv("MINX_DATABASE_URL")),
	}
	if cfg.ServiceName == "" {
		cfg.ServiceName = "minx-core-api"
	}
	if cfg.HTTPAddr == "" {
		cfg.HTTPAddr = ":8080"
	}
	if err := validateHTTPAddr(cfg.HTTPAddr); err != nil {
		return Config{}, fmt.Errorf("invalid MINX_HTTP_ADDR: %w", err)
	}
	return cfg, nil
}

func validateHTTPAddr(addr string) error {
	_, port, err := net.SplitHostPort(addr)
	if err != nil {
		return fmt.Errorf("must be a host:port address: %w", err)
	}
	portNumber, err := strconv.Atoi(port)
	if err != nil || portNumber < 1 || portNumber > 65535 {
		return fmt.Errorf("port must be an integer between 1 and 65535")
	}
	return nil
}
