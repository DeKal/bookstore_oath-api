package users

// User is a user
type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

// LoginRequest is a request to login
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
