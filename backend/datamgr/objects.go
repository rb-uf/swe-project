package datamgr

import "gorm.io/gorm"

type Subject struct {
	gorm.Model
	Name string
}

type Review struct {
	gorm.Model
	Subject Subject
	Rating  uint
	Text    string
	Author	uint
}

type User struct {
	gorm.Model
	ID       uint
	Name     string
	Password string
}
