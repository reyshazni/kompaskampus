package entity

type LectureSubject struct {
	ID              uint          `gorm:"primarykey"`
	SubjectEntityID uint          `gorm:"foreignkey:ID"`
	SubjectEntity   SubjectEntity `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	LectureID       uint          `gorm:"foreignkey:ID"`
	Lecture         Lecture       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
