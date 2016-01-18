package controller

import (
	"github.com/jinzhu/gorm"
	"github.com/syo-sa1982/SeacherServerSide/model"
)

type Controller struct {
	db          gorm.DB
}

type CharaMakeAPI struct {
	BaseStatus  model.PlayerBase
	DiceHistory map[string][]int
}

type SkillSetAPI struct {
	SkillMaster    []model.SkillMaster
	JobMaster      model.JobMaster
	JobSkillMaster []model.JobSkillMaster
	PlayerBase     model.PlayerBase
	PlayerStatus   model.PlayerStatus
}


type UserInfoAPI struct {
	User           model.User
	Job            model.JobMaster
	PlayerBase     model.PlayerBase
	PlayerStatus   model.PlayerStatus
}

type JobSelectAPI struct {
	JobMaster      []model.JobMaster
	JobSkillMaster []model.JobSkillMaster
}

func AppContext(db gorm.DB) Controller {
	return Controller{db: db}
}
