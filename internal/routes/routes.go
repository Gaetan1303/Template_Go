package routes

import (
	"net/http"

	"github.com/Gaetan1303/Template_Go/internal/controllers"
	"github.com/go-chi/chi/v5"
)

// SetupRouter configure les routes principales et prépare l'ajout de services futurs
func SetupRouter() http.Handler {
	r := chi.NewRouter()

	// Middleware global exemple (logging, recovery, etc.)
	// r.Use(middleware.Logger)

	// Route d'accueil
	r.Get("/", controllers.HomeHandler)

	// Prépare l'ajout de routes versionnées ou de groupes de services
	r.Route("/api", func(r chi.Router) {
		r.Mount("/users", userRoutes())
	})

	return r
}
