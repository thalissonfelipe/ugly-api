package models

// User Model
type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// UserResponse model returned on get requests
type UserResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

// Login Model
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
