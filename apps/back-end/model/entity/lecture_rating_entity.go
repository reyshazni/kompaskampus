package entity

import "gorm.io/gorm"

type LectureRating struct {
	gorm.Model
	UserID           uint           `gorm:"foreignkey:ID"`
	User             User           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	LectureSubjectID uint           `gorm:"foreignkey:ID"`
	LectureSubject   LectureSubject `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	IsAnon           bool           `gorm:"type:BOOLEAN;default:false"`
	Message          string         `gorm:"type:TEXT"`
	Recommended      bool           `gorm:"type:BOOLEAN;default:true"`
	Quality          int            `gorm:"DEFAULT 0"`
	Difficulty       int            `gorm:"DEFAULT 0"`
	Overall          float32        `gorm:"DEFAULT 0"`
}
