package model
import "time"

type PlayerBase struct {
	ID int64
	UserID int64 `sql:"not NULL;unique"`
	Name string `sql:"size:255"`

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

type PlayerStatus struct {
	ID       int64
	PlayerID int64 `sql:"not NULL;unique"`
	UserID   int64 `sql:"not NULL;"`

	HP              int
	MP              int
	Sanity          int
	Luck            int
	Idea            int
	Knowledge       int

	JopSkillPoint   int
	HobbySkillPoint int
	DamageBonus     int

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type PlayerSkills struct {
	ID       int64
	PlayerID int64`sql:"not NULL;"`
	SkillValue int
}