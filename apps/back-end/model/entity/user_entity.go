package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Email        string    `gorm:"type:varchar(100);unique;NOT NULL"`
	Username     string    `gorm:"type:varchar(100);unique; NOT NULL"`
	AccessLevel  uint      `gorm:"type:int; DEFAULT:0"`
	FullName     string    `gorm:"type:varchar(100); NOT NULL"`
	Password     string    `gorm:"type:varchar(100); NOT NULL"`
	Avatar       string    `gorm:"type:varchar(100);DEFAULT:''"`
	BirthDate    time.Time `gorm:"type:date; DEFAULT:'1970-01-01'"`
	UniversityID int       //`gorm:"foreignkey:ID"`
	// University   University `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Npm        string `gorm:"type:varchar(50);NOT NULL;unique"`
	IsVerified bool   `gorm:"type:BOOLEAN;default:false"`
}
