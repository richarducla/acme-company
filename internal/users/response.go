package users

import "acme/internal/models"

const (
	statusSuccess = "success"
	statusLogout  = "user logout"
)

type RegisterResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Status   string `json:"status"`
}

type Message struct {
	Message string `json:"message"`
}

func toRegisterResponse(user models.User) RegisterResponse {
	return RegisterResponse{
		Username: user.Username,
		Email:    user.Email,
		Status:   statusSuccess,
	}
}
