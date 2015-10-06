package model
import (
	"github.com/jinzhu/gorm"
)

type SkillMaster struct {
	gorm.Model
	CategoryID int
	SkillName  string
	Value      int
}
type SkillCategory struct {
	gorm.Model
	CategoryName string
}


var SkillCategoryList = []SkillCategory{
	SkillCategory{CategoryName:"治癒"},
	SkillCategory{CategoryName:"戦闘"},
	SkillCategory{CategoryName:"探索"},
}
var SkillMasterList = []SkillMaster{
	SkillMaster{CategoryID:3,SkillName:"言いくるめ",Value:5},
	SkillMaster{CategoryID:1,SkillName:"医学",Value:5},
	SkillMaster{CategoryID:2,SkillName:"回避",Value:0},
}
