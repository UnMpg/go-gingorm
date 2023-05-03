package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID               uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name             string    `gorm:"type:varchar(255);not null "`
	Email            string    `gorm:"uniqueIndex;not null"`
	Password         string    `gorm:"not null"`
	Role             string    `gorm:"type:varchar(255);not null"`
	Provider         string    `gorm:"not null"`
	Photo            string    `gorm:"not null"`
	VerificationCode string
	Verified         bool `gorm:"not null"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type SignUpInput struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
	Photo           string `json:"photo" binding:"required"`
}
type SignInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Name      string    `json:",omitempty"`
	Email     string    `json:",omitempty"`
	Role      string    `json:",omitempty"`
	Photo     string    `json:"photo,omitempty"`
	Provider  string    `json:"provider"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}