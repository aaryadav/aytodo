package middleware

import (
	"net/http"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Add your authentication logic here
		// For example, check for a valid token in the request header
		token := r.Header.Get("Authorization")
		if token != "your-valid-token" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// If the authentication is successful, call the next handler
		next.ServeHTTP(w, r)
	}
}
