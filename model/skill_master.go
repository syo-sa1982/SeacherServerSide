package model
import "github.com/jinzhu/gorm"

type SkillMaster struct {
	gorm.Model
	SkillName string
}