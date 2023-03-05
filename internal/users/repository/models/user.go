package models

import (
	"acme/internal/models"
	"time"
)

type User struct {
	ID        uint      `gorm:"column:id"`
	Username  string    `gorm:"column:username"`
	Password  string    `gorm:"column:password"`
	Email     string    `gorm:"column:email"`
	Status    bool      `gorm:"column:status"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func ToUserDAO(user models.User) User {
	return User{
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
	}
}

func ToUserModel(user User) models.User {
	return models.User{
		ID:        user.ID,
		Username:  user.Username,
		Password:  user.Password,
		Email:     user.Email,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
