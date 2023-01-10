package entity

import "github.com/jinzhu/gorm"

type Lecture struct {
	gorm.Model
	FullName   string `gorm:"type:varchar(100); NOT NULL"`
	University string `gorm:"type:varchar(50);NOT NULL"`
	Faculty    string `gorm:"type:varchar(30); NOT NULL"`
}
