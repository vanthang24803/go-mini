package dto

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type RegisterRequest struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
	Username  string `json:"username" validate:"required"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}

type LogoutRequest struct {
	ID string `json:"id"`
}

type VerifyEmailRequest struct {
	Email string `json:"email" validate:"required,email"`
	Token string `json:"token" validate:"required"`
}

type TokenResponse struct {
	Type         string `json:"type"`
	Exp          int64  `json:"exp"`
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}
