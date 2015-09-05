package model

import "time"

type User struct {
	ID int64
	UUID string `sql:"not NULL;size:36;unique"`
	Name string `sql:"not NULL;size:255;"`
	RollCount int `sql:"DEFAULT:0"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}