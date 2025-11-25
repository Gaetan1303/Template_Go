package middlewares

import (
	"net/http"
)

// Exemple de middleware d'authentification
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Logique d'authentification ici
		next.ServeHTTP(w, r)
	})
}
