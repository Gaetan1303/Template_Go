package factory

import "github.com/Gaetan1303/Template_Go/internal/models"

// UserFactory applique le pattern Factory (créationnel)
// SRP (Single Responsibility Principle) : cette factory ne crée que des Users
// OCP (Open/Closed Principle) : on peut étendre la création sans modifier la factory
func NewUser(id int, name, email string) *models.User {
	return &models.User{ID: id, Name: name, Email: email}
}
