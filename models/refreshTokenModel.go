package models

import "github.com/google/uuid"

type RefreshToken struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Token     *string    `gorm:"not null;unique" json:"token"`
	ExpiresAt *int64     `gorm:"not null" json:"expired_at"`
	UserID    *uuid.UUID `gorm:"type:uuid;not null;" json:"user_id"`
	User      User        `gorm:"foreignKey:UserID;" json:"-"`
}