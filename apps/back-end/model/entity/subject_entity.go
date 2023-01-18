package entity

import "gorm.io/gorm"

type Subject struct {
	gorm.Model
	Name         string `gorm:"type:varchar(100); NOT NULL"`
	UniversityID int    //`gorm:"foreignkey:ID"`
	//University   University `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Description string `gorm:"type:TEXT"`
}
