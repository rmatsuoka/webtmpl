package xsql

import (
	"context"
	"database/sql"
)

func ExecLastInsertID(ctx context.Context, tx *sql.Tx, query string, args ...any) (id int64, err error) {
	r, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	return r.LastInsertId()
}

func ExecRowsAffected(ctx context.Context, tx *sql.Tx, query string, args ...any) (rows int64, err error) {
	r, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}
