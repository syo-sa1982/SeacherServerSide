package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model

	UUID string `sql:"not NULL;size:36;unique"`
	Name string `sql:"size:255;"`
	RollCount int `sql:"DEFAULT:0"`
}