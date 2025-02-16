package context

import (
	"context"
	"lukachi/eth-indexer/internal/db"
	"net/http"
)

type ctxKey int

const (
	DBCtxKey ctxKey = iota
)

func CtxMiddleWare(extenders ...func(context.Context) context.Context) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			for _, extender := range extenders {
				ctx = extender(ctx)
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func CtxDB(entry db.DB) func(ctx context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, DBCtxKey, entry)
	}
}
