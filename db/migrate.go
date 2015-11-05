package main
import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/syo-sa1982/SeacherServerSide/model"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"os"
	"log"
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

	log.Println("User")
	if db.HasTable(&model.User{}) { db.DropTable(&model.User{}) }
	db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.User{})

	log.Println("PlayerBase")
	if db.HasTable(&model.PlayerBase{}) { db.DropTable(&model.PlayerBase{}) }
	db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.PlayerBase{})

	log.Println("PlayerStatus")
	if db.HasTable(&model.PlayerStatus{}) { db.DropTable(&model.PlayerStatus{}) }
	db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.PlayerStatus{})

	log.Println("PlayerSkill")
	if db.HasTable(&model.PlayerSkill{}) { db.DropTable(&model.PlayerSkill{}) }
	db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.PlayerSkill{})

	log.Println("SkillCategory")
	if db.HasTable(&model.SkillCategory{}) { db.DropTable(&model.SkillCategory{}) }
	db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.SkillCategory{})

	log.Println("SkillCategory INSERT")
	for key := range model.SkillCategoryList{
		db.Create(&model.SkillCategoryList[key])
	}

	log.Println("SkillMaster")
	if db.HasTable(&model.SkillMaster{}) { db.DropTable(&model.SkillMaster{}) }
	db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.SkillMaster{})

	log.Println("SkillMaster INSERT")
	for key := range model.SkillMasterList{
		db.Create(&model.SkillMasterList[key])
	}

}

