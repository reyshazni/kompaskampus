package entity

import "github.com/jinzhu/gorm"

type University struct {
	gorm.Model
	FullName string `gorm:"type:varchar(100); NOT NULL"`
	UniCode  string `gorm:"type:varchar(10); NOT NULL"`
}
