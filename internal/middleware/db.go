package middleware

import (
	"database/sql"
	"net/http"

	"github.com/ekozlova94/parking/internal/ctxutils"
)

func Db(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		withContext := r.WithContext(ctxutils.NewDbContext(r.Context(), db))
		next.ServeHTTP(w, withContext)
	})
}
