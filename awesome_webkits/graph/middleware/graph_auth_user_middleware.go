package middleware

import (
	"awesome_webkits/database/models"
	"awesome_webkits/utils/parser"
	"context"
	"net/http"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func GraphQLAuthUserMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenHeader := r.Header.Get("Authorization")

			// Allow unauthenticated users in
			if tokenHeader == "" {
				next.ServeHTTP(w, r)
				return
			}

			token, err := parser.ReadBearerToken(tokenHeader)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			userID, validErr := models.OauthAccessToken{}.VerifyToken(token)
			if validErr != nil {
				http.Error(w, validErr.Message, validErr.Status)
				return
			}

			user := models.User{}
			user.ID = userID
			if err := user.FindWithId(); err != nil {
				http.Error(w, err.Message, err.Status)
				return
			}
			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, &user)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.sh.
func GraphQLForContext(ctx context.Context) *models.User {
	raw, _ := ctx.Value(userCtxKey).(*models.User)
	return raw
}
