package interfaces

// Interface for auth service
type IAuthService interface {
	HashPassword(password string) ([]byte, error)
	ComparePasswords(password string, hashedPassword string) error
}
