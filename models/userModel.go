package models

// User defines the user model
// @Description A user represents the user details.
// @Type object
type User struct {
	ID        string `json:"id"`
	Firstname string `json:"first" `
	Lastname  string `json:"last"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=8"`
	Role      string `json:"role"`
}

// ErrorResponse represents the error response structure
// @Description This is the structure for error responses.
// @Type object
type ErrorResponse struct {
	Error string `json:"error"`
}
