package postgres

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrMissingDatabaseURL = errors.New("MINX_DATABASE_URL is required")

type PoolSettings struct {
	MaxConns        int32
	MinConns        int32
	MaxConnLifetime time.Duration
	MaxConnIdleTime time.Duration
}

func DefaultPoolSettings() PoolSettings {
	return PoolSettings{
		MaxConns:        20,
		MinConns:        2,
		MaxConnLifetime: time.Hour,
		MaxConnIdleTime: 30 * time.Minute,
	}
}

func NewPool(ctx context.Context, databaseURL string, settings PoolSettings) (*pgxpool.Pool, error) {
	if databaseURL == "" {
		return nil, ErrMissingDatabaseURL
	}
	if settings.MaxConns <= 0 || settings.MinConns < 0 || settings.MinConns > settings.MaxConns {
		return nil, errors.New("invalid PostgreSQL pool connection limits")
	}
	if settings.MaxConnLifetime < 0 || settings.MaxConnIdleTime < 0 {
		return nil, errors.New("PostgreSQL pool lifetimes cannot be negative")
	}

	cfg, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		return nil, fmt.Errorf("parse PostgreSQL connection URL: %w", err)
	}
	cfg.MaxConns = settings.MaxConns
	cfg.MinConns = settings.MinConns
	cfg.MaxConnLifetime = settings.MaxConnLifetime
	cfg.MaxConnIdleTime = settings.MaxConnIdleTime
	cfg.ConnConfig.RuntimeParams["application_name"] = "minx-core-api"

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("create PostgreSQL pool: %w", err)
	}

	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := pool.Ping(pingCtx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("verify PostgreSQL connection: %w", err)
	}
	return pool, nil
}
