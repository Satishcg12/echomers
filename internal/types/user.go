package types

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	FullName          string     `json:"full_name"`
	Username          string     `json:"username"`
	Email             string     `json:"email"`
	Password          string     `json:"password"`
	VerificationToken string     `json:"verification_token"`
	VerifiedAt        *time.Time `json:"verified_at"`
}
