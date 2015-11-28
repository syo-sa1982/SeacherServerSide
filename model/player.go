package model

import (
	"time"
)

type PlayerBase struct {
	ID int

	UserID int   `sql:"not NULL;"`
	Name   string `sql:"size:255"`

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
	ID int

	UserID   int `sql:"not NULL;"`
	PlayerID int `sql:"not NULL;unique_index"`
	JobID    int

	MaxHP     int
	MaxMP     int
	HP        int
	MP        int
	Sanity    int
	Luck      int
	Idea      int
	Knowledge int

	JobSkillPoint   int
	HobbySkillPoint int
	DamageBonus     int

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type PlayerSkill struct {
	ID int

	PlayerID int `sql:"not NULL;index"`
	SkillID  int `sql:"not NULL;index"`
	Value    int
}
