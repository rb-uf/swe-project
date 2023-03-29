package datamgr

import "gorm.io/gorm"

type Subject struct {
	gorm.Model
	Name string
}

type Review struct {
	gorm.Model
	Subject  string
	Rating   uint
	Text     string
	Author   string
	AuthorID uint
}

type User struct {
	gorm.Model

	Name     string
	Hash     []byte

	Admin bool
}
