package models

type User struct {
	ID int64
	UserID int64
	Name string `sql:"size:255"`
	HP int
	San int

}