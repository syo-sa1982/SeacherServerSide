package main
import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/syo-sa1982/SeacherServerSide/model"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"os"
)


func main(){
	yml, err := ioutil.ReadFile("conf/db.yaml")
	if err != nil {
		panic(err)
	}

	t := make(map[interface{}]interface{})

	_ = yaml.Unmarshal([]byte(yml), &t)

	conn := t[os.Getenv("GOJIENV")].(map[interface{}]interface{})

	db, err := gorm.Open("mysql", conn["user"].(string)+conn["password"].(string)+"@/"+conn["db"].(string)+"?charset=utf8&parseTime=True")
	if err != nil {
		panic(err)
	}

	if db.HasTable(&model.User{}) { db.DropTable(&model.User{}) }
	db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.User{})

	if db.HasTable(&model.PlayerBase{}) { db.DropTable(&model.PlayerBase{}) }
	db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.PlayerBase{})

	if db.HasTable(&model.PlayerStatus{}) { db.DropTable(&model.PlayerStatus{}) }
	db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.PlayerStatus{})

	if db.HasTable(&model.PlayerSkill{}) { db.DropTable(&model.PlayerSkill{}) }
	db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.PlayerSkill{})

	if db.HasTable(&model.SkillCategory{}) { db.DropTable(&model.SkillCategory{}) }
	db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.SkillCategory{})

	for key := range model.SkillCategoryList{
		db.Create(&model.SkillCategoryList[key])
	}

	if db.HasTable(&model.SkillMaster{}) { db.DropTable(&model.SkillMaster{}) }
	db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.SkillMaster{})

	for key := range model.SkillMasterList{
		db.Create(&model.SkillMasterList[key])
	}

}

