package xsql

import (
	"cmp"
	"context"
	"database/sql"
	"iter"
)

type Querier interface {
	QueryRowContext(context.Context, string, ...any) *sql.Row
	QueryContext(context.Context, string, ...any) (*sql.Rows, error)
}

type Rows struct {
	rows   *sql.Rows
	err    error
	closed bool
}

func (r *Rows) Scan() iter.Seq[func(...any)] {
	return func(yield func(func(...any)) bool) {
		if r.closed || r.err != nil {
			return
		}
		defer func() {
			r.rows.Close()
			r.closed = true
		}()
		for r.rows.Next() {
			if !yield(func(dest ...any) { r.err = r.rows.Scan(dest...) }) {
				return
			}

			if r.err != nil {
				return
			}
		}
	}
}

func (r *Rows) Err() error {
	return cmp.Or(r.err, r.rows.Err())
}

func Query(ctx context.Context, db Querier, query string, args ...any) *Rows {
	rows := new(Rows)
	rows.rows, rows.err = db.QueryContext(ctx, query, args...)
	return rows
}
