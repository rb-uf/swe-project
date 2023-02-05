// Struct definition for a user

package objects

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int
	Name     string
	Password string
}
