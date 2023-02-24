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
}

type User struct {
	gorm.Model
	ID       int
	Name     string
	Password string
}
