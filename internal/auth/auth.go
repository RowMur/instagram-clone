package auth

import (
	"context"
	"github/rowmur/insta-clone/internal/database"
	"net/http"
	"strings"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func Middleware(dbQueries *database.Queries) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apiKey := strings.Replace(r.Header.Get("Authorization"), "ApiKey ", "", 1)
			if apiKey == "" {
				h.ServeHTTP(w, r)
				return
			}

			user, err := dbQueries.GetUserByApiKey(context.Background(), apiKey)
			if err != nil {
				http.Error(w, "Invalid API key", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), userCtxKey, &user)
			r = r.WithContext(ctx)
			h.ServeHTTP(w, r)
		})
	}
}

func ForContext(ctx context.Context) *database.User {
	raw, _ := ctx.Value(userCtxKey).(*database.User)
	return raw
}
