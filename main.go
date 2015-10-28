package main

import (
	"log"
	"github.com/zenazn/goji"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
	"github.com/syo-sa1982/SeacherServerSide/controller"
	"os"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var (
	db gorm.DB
	cntr controller.Controller
)

func main() {
	log.Print("main")
	user := web.New()
	player := web.New()

	goji.Handle("/user/*", user)
	goji.Handle("/player/*", player)

	user.Use(middleware.SubRouter)
	user.Get("/index", cntr.UserIndex)
	user.Post("/add", cntr.UserAdd)
	user.Post("/auth", cntr.UserAuth)

	player.Use(middleware.SubRouter)
	player.Post("/base_make", cntr.PlayerBaseMake)
	player.Post("/generate", cntr.PlayerGenerate)
	player.Post("/list", cntr.PlayerList)
	player.Get("/skill_setting", cntr.SkillSetting)
	player.Post("/skill_setting", cntr.SkillSetting)
	goji.Serve()
}

func init() {
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
	cntr = controller.AppContext(db)
}
