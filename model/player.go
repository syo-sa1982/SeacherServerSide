package model
import (
	"github.com/jinzhu/gorm"
)

type PlayerBase struct {
	gorm.Model

	UserID uint `sql:"not NULL;"`
	Name string `sql:"size:255"`

	Strength     int
	Constitution int
	Power        int
	Dextality    int
	Appeal       int
	Size         int
	Intelligence int
	Education    int
}

type PlayerStatus struct {
	gorm.Model

	UserID   uint `sql:"not NULL;"`
	PlayerID uint`sql:"not NULL;index"`

	MaxHP           int
	MaxMP           int
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
	PlayerID uint`sql:"not NULL;index;"`
	SkillIDs string
	SkillValue int
}