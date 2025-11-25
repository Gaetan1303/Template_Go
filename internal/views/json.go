package views

import (
	"encoding/json"
	"net/http"
)

// RenderJSON applique SRP et OCP (une seule responsabilité, extensible)
func RenderJSON(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}
