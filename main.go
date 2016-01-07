package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/syo-sa1982/SeacherServerSide/controller"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
//	"github.com/zenazn/goji/web/middleware"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"net"
	"net/http/fcgi"
	"net/http"
	"github.com/syo-sa1982/SeacherServerSide/admin"
)

var (
	db   gorm.DB
	cntr controller.Controller
	ad admin.Controller
)

func rooter(m *web.Mux) http.Handler {

	m.Get("/admin/", ad.AdminIndex)

	m.Get("/user/index", cntr.UserIndex)
	m.Post("/user/add", cntr.UserAdd)
	m.Post("/user/auth", cntr.UserAuth)

	m.Get("/player/joblist", cntr.JobList)
	m.Post("/player/joblist", cntr.JobList)
	m.Post("/player/base_make", cntr.PlayerBaseMake)
	m.Post("/player/generate", cntr.PlayerGenerate)
	m.Post("/player/skill_setting", cntr.SkillSetting)
	m.Post("/player/skill_submit", cntr.SkillSubmit)

	m.Post("/home/user/info", cntr.UserInfo)
	m.Get("/home/scenario/list", cntr.ScenarioList)
	m.Post("/home/player/list", cntr.PlayerList)

	return m
}

func main() {
	log.Print("main")

	listener,_ := net.Listen("tcp", ":9000")
	fcgi.Serve(listener, rooter(goji.DefaultMux))

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
	ad = admin.AppContext(db)
}
