package main
import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/syo-sa1982/SeacherServerSide/model"
)


func main(){
	db, _ := gorm.Open("mysql", "root:@/seacher?charset=utf8&parseTime=True")
	if !db.HasTable(&model.User{}) {
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.User{})
	}
	if !db.HasTable(&model.PlayerBase{}) {
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.PlayerBase{})
	}
	if !db.HasTable(&model.PlayerStatus{}) {
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.PlayerStatus{})
	}
	if !db.HasTable(&model.PlayerSkill{}) {
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.PlayerSkill{})
	}

	if db.HasTable(&model.SkillCategory{}) {
		db.DropTable(&model.SkillCategory{})
	}
	if !db.HasTable(&model.SkillCategory{}) {
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.SkillCategory{})
	}
	for key := range model.SkillCategoryList{
		db.Create(&model.SkillCategoryList[key])
	}

	if db.HasTable(&model.SkillMaster{}) {
		db.DropTable(&model.SkillMaster{})
	}
	db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.SkillMaster{})

	for key := range model.SkillMasterList{
		db.Create(&model.SkillMasterList[key])
	}

}

