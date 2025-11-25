package routes

import (
	"github.com/Gaetan1303/Template_Go/internal/controllers"
	"github.com/go-chi/chi/v5"
)

// userRoutes prépare les routes utilisateurs, facilement extensible
func userRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", controllers.HomeHandler) // exemple, à remplacer par UserController
	// r.Post("/", controllers.CreateUser)
	return r
}
