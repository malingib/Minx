package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrInvalidAccountID = errors.New("invalid authenticated account ID")
	ErrNilTransactionFn = errors.New("account transaction callback is required")
)

// InAccountTx binds the verified actor to the PostgreSQL transaction so RLS
// policies can enforce account and tenant visibility for every query.
func InAccountTx(ctx context.Context, pool *pgxpool.Pool, accountID string, fn func(pgx.Tx) error) error {
	if fn == nil {
		return ErrNilTransactionFn
	}
	var actor pgtype.UUID
	if err := actor.Scan(accountID); err != nil || !actor.Valid {
		return ErrInvalidAccountID
	}
	if pool == nil {
		return errors.New("PostgreSQL pool is required")
	}

	tx, err := pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin account transaction: %w", err)
	}
	defer func() {
		_ = tx.Rollback(context.Background())
	}()

	if _, err := tx.Exec(ctx, "SELECT set_config('minx.account_id', $1, true)", accountID); err != nil {
		return fmt.Errorf("set authenticated account context: %w", err)
	}
	if err := fn(tx); err != nil {
		return err
	}
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit account transaction: %w", err)
	}
	return nil
}
