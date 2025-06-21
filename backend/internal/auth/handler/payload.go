package handler

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
}

type RegisterResponse struct {
	Token string `json:"token"`
}

const (
	ErrInternalServer       = "Internal server error"
	ErrInvalidRequestBody   = "Invalid request body"
	ErrIncorrectСredentials = "Incorrect credentials"
)
