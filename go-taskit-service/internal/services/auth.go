package service

import "golang.org/x/crypto/bcrypt"

// Represents a auth service with functionality for hashing and comparing passwords
type AuthService struct {
}

// Constructor function
func NewAuthService() *AuthService {
	return &AuthService{}
}

// Function for hashing password
func (a *AuthService) HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// Function for comparing a plain and hashed password for equality
func (a *AuthService) ComparePasswords(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
