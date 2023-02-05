// Struct definition for a post

package objects

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title string
	Tags  []string
	Score int
	Body  string

	Likes int
}
