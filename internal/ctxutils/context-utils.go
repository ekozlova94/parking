package ctxutils

import (
	"context"
	"database/sql"
)

func NewDbContext(ctx context.Context, db *sql.DB) context.Context {
	return context.WithValue(ctx, dbKey, db)
}

func DbFromContext(ctx context.Context) *sql.DB {
	return ctx.Value(dbKey).(*sql.DB)
}
