package model

import "gorm.io/gorm"

type TestModel struct {
	gorm.Model
	Text string
}
