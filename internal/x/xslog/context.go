package xslog

import (
	"context"
)

func ContextWith(ctx context.Context, args ...any) context.Context {
	v, _ := ctx.Value(argsKey).([][]any)
	newv := make([][]any, len(v)+1)
	copy(newv, v)
	newv[len(v)] = args
	return context.WithValue(ctx, argsKey, newv)
}
