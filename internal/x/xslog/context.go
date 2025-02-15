package xslog

import (
	"context"
)

type contextArgs [][]any

func ContextWith(ctx context.Context, args ...any) context.Context {
	v, _ := ctx.Value(attrKey).(contextArgs)
	newv := make(contextArgs, len(v)+1)
	copy(newv, v)
	newv[len(v)] = args
	return context.WithValue(ctx, attrKey, newv)
}
