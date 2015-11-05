package controller
import (
	"github.com/jinzhu/gorm"
	"github.com/syo-sa1982/SeacherServerSide/model"
)

type Controller struct {
	db gorm.DB
	charaStatus CharaMakeAPI
	SkillSetAPI
}

type CharaMakeAPI struct {
	BaseStatus, CharaStatus map[string]int
	DiceHistory map[string][]int
}


type SkillSetAPI struct {
	SkillMaster    []model.SkillMaster
	JobMaster      []model.JobMaster
	JobSkillMaster []model.JobSkillMaster
	PlayerStatus   model.PlayerStatus
}

func AppContext(db gorm.DB) Controller {
	return Controller{db: db}
}

