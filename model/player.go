package model
import (
	"github.com/jinzhu/gorm"
	"database/sql"
)

type PlayerBase struct {
	gorm.Model

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

	Status PlayerStatus
	StatusID sql.NullInt64
	PlayerSkills []PlayerSkill
}

type PlayerStatus struct {
	gorm.Model

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
}

type PlayerSkill struct {
	gorm.Model

	PlayerID int64`sql:"not NULL;"`
	SkillValue int
}