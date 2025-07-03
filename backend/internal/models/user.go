package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BasicModel struct {
	Id        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type User struct {
	BasicModel
	Name     string `gorm:"not null"`
	Age      int    `gorm:"not null"`
	CodeHash string `gorm:"not null"`
}

// BeforeCreate hook to auto-generate UUID if not set
func (b *BasicModel) BeforeCreate(tx *gorm.DB) error {
	if b.Id == uuid.Nil {
		b.Id = uuid.New()
	}
	return nil
}
