package models

import (
	"time"

	"gorm.io/gorm"
)

type RevokedToken struct {
	gorm.Model
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}
