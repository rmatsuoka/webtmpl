package xsql

import (
	"context"
	"database/sql"
	"errors"
)

func ExecLastInsertID(ctx context.Context, tx *sql.Tx, query string, args ...any) (id int64, err error) {
	r, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	id, err = r.LastInsertId()
	if errors.Is(err, errors.ErrUnsupported) {
		panic("LastInsertId is not supported by the driver")
	}
	return id, err
}

func ExecRowsAffected(ctx context.Context, tx *sql.Tx, query string, args ...any) (rows int64, err error) {
	r, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	rows, err = r.RowsAffected()
	if errors.Is(err, errors.ErrUnsupported) {
		panic("RowsAffected is not supported by the driver")
	}
	return rows, err
}
