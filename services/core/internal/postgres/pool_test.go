package postgres

import (
	"context"
	"errors"
	"testing"

	"github.com/jackc/pgx/v5"
)

func TestDefaultPoolSettings(t *testing.T) {
	got := DefaultPoolSettings()
	if got.MaxConns != 20 || got.MinConns != 2 {
		t.Fatalf("unexpected pool limits: %+v", got)
	}
}

func TestNewPoolRequiresDatabaseURL(t *testing.T) {
	_, err := NewPool(context.Background(), "", DefaultPoolSettings())
	if !errors.Is(err, ErrMissingDatabaseURL) {
		t.Fatalf("expected missing database URL, got %v", err)
	}
}

func TestNewPoolRejectsInvalidLimits(t *testing.T) {
	settings := DefaultPoolSettings()
	settings.MinConns = settings.MaxConns + 1
	_, err := NewPool(context.Background(), "postgres://user:pass@localhost:5432/minx", settings)
	if err == nil {
		t.Fatal("expected invalid pool limits")
	}
}

func TestInAccountTxRejectsInvalidActorBeforeUsingPool(t *testing.T) {
	err := InAccountTx(context.Background(), nil, "not-a-uuid", func(_ pgx.Tx) error {
		return nil
	})
	if !errors.Is(err, ErrInvalidAccountID) {
		t.Fatalf("expected invalid account ID, got %v", err)
	}
}

func TestInAccountTxRequiresCallback(t *testing.T) {
	err := InAccountTx(context.Background(), nil, "10000000-0000-4000-8000-000000000001", nil)
	if !errors.Is(err, ErrNilTransactionFn) {
		t.Fatalf("expected missing callback, got %v", err)
	}
}

func TestInAccountTxValidatesUUIDBeforeRequiringPool(t *testing.T) {
	err := InAccountTx(context.Background(), nil, "10000000-0000-4000-8000-000000000001", func(_ pgx.Tx) error {
		return nil
	})
	if err == nil || err.Error() != "PostgreSQL pool is required" {
		t.Fatalf("expected missing pool error after valid UUID, got %v", err)
	}
}
