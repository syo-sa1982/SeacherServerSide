package model

import (
	"time"
)

type PlayerBase struct {
	ID uint

	UserID uint   `sql:"not NULL;"`
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
	ID uint

	UserID   uint `sql:"not NULL;"`
	PlayerID uint `sql:"not NULL;unique_index"`
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
	ID uint

	PlayerID uint `sql:"not NULL;index"`
	SkillID  uint `sql:"not NULL;index"`
	Value    int
}
