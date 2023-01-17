package entity

import (
	"gorm.io/gorm"
	"time"
)

type RefreshToken struct {
	UserID     uint `gorm:"primarykey"`
	User       User
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	RefreshKey string
}
