package controllers

import (
	"net/http"

	"github.com/Gaetan1303/Template_Go/internal/views"
)

// HomeHandler = Controller (MVC)
// SRP : gère uniquement la requête d'accueil
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Utilisation de la vue (View) pour la réponse
	views.RenderJSON(w, map[string]string{"message": "Bienvenue sur l'API Go!"}, http.StatusOK)
}
