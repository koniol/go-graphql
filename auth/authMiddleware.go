package auth

import (
	"context"
	"github.com/go-pg/pg/v10"
	"net/http"
)

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var userCtxKey = &contextKey{"user"}
type contextKey struct {
	email string
}

// Middleware decodes the share session cookie and packs the session into context
func Middleware(db *pg.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			// Allow unauthenticated users in
			if token == "" {
				//http.Error(w, "Unauthenticated user please login", http.StatusForbidden)
				next.ServeHTTP(w, r)
				return
			}

			user, err := GetTokenData(token)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusForbidden)
				return
			}

			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, user)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *UserToken {
	raw, _ := ctx.Value(userCtxKey).(*UserToken)
	return raw
}
