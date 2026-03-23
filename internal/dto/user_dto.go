package dto

type RegisterRequest struct {
    Name string `json:"name"`
    Email string `json:"email" binding:"required,email"`
    Phone string `json:"phone"`
    Password string `json:"password" binding:"required,min=6"`
    PasswordConfirm string `json:"password_confirm" binding:"required,min=6"`
}

type LoginRequest struct {
    Email string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	Id uint `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type RegisterResponse struct {
	User UserResponse `json:"user"`
}

type LoginResponse struct {
	User UserResponse `json:"user"`
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type RefreshResponse struct {
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}