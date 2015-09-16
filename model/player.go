package model
import "time"

type Player struct {
	ID int64
	UserID int64 `sql:"not NULL;unique"`
	Name string `sql:"size:255"`
	HP           int
	Sanity       int
	Strength     int
	Constitution int
	Power        int
	Dextality    int
	Appeal       int
	Size         int
	Intelligence int
	Education    int

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}