package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/syo-sa1982/SeacherServerSide/model"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

func main() {
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

	if os.Args[1] == "setup" {
		log.Println("User")
		if db.HasTable(&model.User{}) {
			db.DropTable(&model.User{})
		}
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.User{})
		log.Println("User Done")

		log.Println("PlayerBase")
		if db.HasTable(&model.PlayerBase{}) {
			db.DropTable(&model.PlayerBase{})
		}
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.PlayerBase{})
		log.Println("PlayerBase Done")

		log.Println("PlayerStatus")
		if db.HasTable(&model.PlayerStatus{}) {
			db.DropTable(&model.PlayerStatus{})
		}
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.PlayerStatus{})
		log.Println("PlayerStatus Done")

		log.Println("PlayerSkill")
		if db.HasTable(&model.PlayerSkill{}) {
			db.DropTable(&model.PlayerSkill{})
		}
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.PlayerSkill{})
		log.Println("PlayerSkill Done")
	}

	log.Println("SkillCategory")
	if db.HasTable(&model.SkillCategory{}) {
		db.DropTable(&model.SkillCategory{})
	}
	db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.SkillCategory{})

	log.Println("SkillCategory INSERT")
	for key := range model.SkillCategoryList {
		db.Create(&model.SkillCategoryList[key])
	}
	log.Println("SkillCategory Done")

	log.Println("SkillMaster")
	if db.HasTable(&model.SkillMaster{}) {
		db.DropTable(&model.SkillMaster{})
	}
	db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.SkillMaster{})

	log.Println("SkillMaster INSERT")
	for key := range model.SkillMasterList {
		db.Create(&model.SkillMasterList[key])
	}
	log.Println("SkillMaster Done")

	log.Println("JobMaster")
	if db.HasTable(&model.JobMaster{}) {
		db.DropTable(&model.JobMaster{})
	}
	db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.JobMaster{})

	log.Println("JobMaster INSERT")
	for key := range model.JobList {
		db.Create(&model.JobList[key])
	}
	log.Println("JobMaster Done")

	log.Println("JobSkillMaster")
	if db.HasTable(&model.JobSkillMaster{}) {
		db.DropTable(&model.JobSkillMaster{})
	}
	db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.JobSkillMaster{})

	log.Println("JobSkillMaster INSERT")
	for key := range model.JobSkillList {
		db.Create(&model.JobSkillList[key])
	}
	log.Println("JobSkillMaster Done")

}
