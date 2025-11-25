package repositories

import "github.com/Gaetan1303/Template_Go/internal/models"

// Exemple de repository utilisateur
type UserRepository struct{}

func (r *UserRepository) FindByID(id int) *models.User {
	return &models.User{ID: id, Name: "John Doe", Email: "john@example.com"}
}
