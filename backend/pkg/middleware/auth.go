package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/vnkot/piklnk/configs"
	"github.com/vnkot/piklnk/pkg/apierr"
	"github.com/vnkot/piklnk/pkg/jwt"
	"github.com/vnkot/piklnk/pkg/res"
)

type key string

const ContextJWTDataKey key = "ContextJWTDataKey"

func writeUnauth(w http.ResponseWriter) {
	res.Json(w, apierr.New(http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized, ""), http.StatusUnauthorized)
}

func IsAuthed(next http.Handler, config configs.AuthConfig) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authedHeader := r.Header.Get("Authorization")

		if !strings.HasPrefix(authedHeader, "Bearer ") {
			writeUnauth(w)
			return
		}

		token := strings.TrimPrefix(authedHeader, "Bearer ")
		jwtData, isValid := jwt.NewJWT(config.Secret).Parse(token)

		if !isValid {
			writeUnauth(w)
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ContextJWTDataKey, jwtData)))
	})
}

func IsMaybeAuthed(next http.Handler, config configs.AuthConfig) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authedHeader := r.Header.Get("Authorization")

		if strings.HasPrefix(authedHeader, "Bearer ") {
			token := strings.TrimPrefix(authedHeader, "Bearer ")
			jwtData, isValid := jwt.NewJWT(config.Secret).Parse(token)

			if isValid {
				next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ContextJWTDataKey, jwtData)))
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
