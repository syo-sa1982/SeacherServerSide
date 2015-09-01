package models

import "time"

type User struct {
	ID int64
	UUID string `sql:"size:36;index"`
	Name string `sql:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}