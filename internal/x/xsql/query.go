package xsql

import (
	"context"
	"database/sql"
	"iter"
)

type Querier interface {
	QueryRowContext(context.Context, string, ...any) *sql.Row
	QueryContext(context.Context, string, ...any) (*sql.Rows, error)
}

func Query(ctx context.Context, db Querier, query string, args ...any) iter.Seq2[func(...any), error] {
	return func(yield func(func(...any), error) bool) {
		rows, err := db.QueryContext(ctx, query, args...)
		if err != nil {
			yield(nil, err)
			return
		}
		defer rows.Close()

		for rows.Next() {
			if !yield(func(dest ...any) { err = rows.Scan(dest...) }, nil) {
				return
			}
			if err != nil {
				yield(nil, err)
				return
			}
		}
		if err := rows.Err(); err != nil {
			yield(nil, err)
		}
	}
}
