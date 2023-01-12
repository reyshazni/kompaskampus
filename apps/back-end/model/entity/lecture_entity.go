package entity

import "gorm.io/gorm"

type Lecture struct {
	gorm.Model
	FullName     string     `gorm:"type:varchar(100); NOT NULL"`
	UniversityID int        `gorm:"foreignkey:ID"`
	University   University `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Faculty      string     `gorm:"type:varchar(100); NOT NULL"`
	Href         string     `gorm:"type:varchar(100)"`
}
