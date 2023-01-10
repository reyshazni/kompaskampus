package entity

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Email      string    `gorm:"type:varchar(100);unique_index;NOT NULL"`
	Username   string    `gorm:"type:varchar(100);unique_index; NOT NULL"`
	FullName   string    `gorm:"type:varchar(100); NOT NULL"`
	Password   string    `gorm:"type:varchar(100); NOT NULL"`
	Avatar     string    `gorm:"type:varchar(100);DEFAULT:''"`
	BirthDate  time.Time `gorm:"type:date; DEFAULT:'1970-01-01'"`
	University string    `gorm:"type:varchar(50);NOT NULL"`
	Npm        string    `gorm:"type:varchar(50);NOT NULL;unique_index"`
	IsVerified bool      `gorm:"type:BOOLEAN;default:false"`
}
