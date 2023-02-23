// Struct definitions

package objects

import "gorm.io/gorm"

type Subject struct {
	gorm.Model
	Name		string
}

type Review struct {
	gorm.Model
	Subject		Subject
	Rating		uint
	Text		string
}

type Post struct {
	gorm.Model
	Title string
	Tags  []string
	Score int
	Body  string

	Likes int
}

type User struct {
	gorm.Model
	ID       int
	Name     string
	Password string
}
