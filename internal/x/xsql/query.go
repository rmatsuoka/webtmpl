package xsql

import (
	"cmp"
	"context"
	"database/sql"
)

type Querier interface {
	QueryContext(context.Context, string, ...any) (*sql.Rows, error)
}

type Rows struct {
	*sql.Rows
	err error
}

func (r *Rows) Scan(dest ...any) {
	if r.err != nil {
		return
	}
	r.err = r.Rows.Scan(dest...)
}

func (r *Rows) Next() bool {
	if r.err != nil {
		return false
	}
	return r.Rows.Next()
}

func (r *Rows) Err() error {
	return cmp.Or(r.err, r.Rows.Err())
}

func (r *Rows) Close() error {
	if r.Rows != nil {
		closeErr := r.Rows.Close()
		return cmp.Or(r.err, closeErr)
	}
	return r.err
}

func Query(ctx context.Context, db Querier, query string, args ...any) *Rows {
	rows := new(Rows)
	rows.Rows, rows.err = db.QueryContext(ctx, query, args...)
	return rows
}
