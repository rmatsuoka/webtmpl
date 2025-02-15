package xsql

import (
	"context"
	"database/sql"
)

func WithTx(ctx context.Context, db *sql.DB, f func(*sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	if err := f(tx); err != nil {
		return err
	}
	return tx.Commit()
}

func WithTx2[T any](ctx context.Context, db *sql.DB, f func(*sql.Tx) (T, error)) (T, error) {
	var t T
	err := WithTx(ctx, db, func(tx *sql.Tx) error {
		var err error
		t, err = f(tx)
		return err
	})
	return t, err
}
