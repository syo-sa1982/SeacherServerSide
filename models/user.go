package models

import "time"

type User struct {
	ID int64
	UUID string `sql:"not NULL;size:36;unique"`
	RollCount int `sql:"DEFAULT:0"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}