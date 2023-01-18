package entity

type University struct {
	ID       int    `gorm:"primary_key" faker:"-"`
	FullName string `gorm:"type:varchar(100); NOT NULL" faker:"first_name"`
	UniCode  string `gorm:"type:varchar(10); NOT NULL" faker:"username"`
}
