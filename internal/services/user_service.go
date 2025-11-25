package services

import "github.com/Gaetan1303/Template_Go/internal/models"

// Exemple de service utilisateur
// UserService = Service (MVC)
// SRP : gère la logique métier liée à l'utilisateur
// OCP : peut être étendu sans modification
// DIP : dépend de l'abstraction (interface UserRepository possible)
type UserService struct{}

func (s *UserService) GetUser(id int) *models.User {
	return &models.User{ID: id, Name: "John Doe", Email: "john@example.com"}
}
