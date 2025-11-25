package models

// User = Model (MVC)
// SRP : représente uniquement la structure d'un utilisateur
type User struct {
	ID    int
	Name  string
	Email string
}
