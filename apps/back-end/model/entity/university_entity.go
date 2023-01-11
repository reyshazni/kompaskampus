package entity

type University struct {
	ID       int    `gorm:"primary_key"`
	FullName string `gorm:"type:varchar(100); NOT NULL"`
	UniCode  string `gorm:"type:varchar(10); NOT NULL"`
}
