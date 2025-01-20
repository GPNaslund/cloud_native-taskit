package dto

// Represents a user
type User struct {
	Username string `json:"username" xml:"username" form:"username" binding:"required"`
	Password string `json:"password" xml:"password" form:"password" binding:"required"`
}
