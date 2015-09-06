package main

import (
	"log"
	"github.com/zenazn/goji"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
	"github.com/syo-sa1982/SeacherServerSide/controller"
)

var (
	db gorm.DB
	cntr controller.Controller
)

func main() {
	log.Print("main")
	user := web.New()

	goji.Handle("/user/*", user)

	user.Use(middleware.SubRouter)
	user.Get("/index", cntr.UserIndex)
	user.Get("/add", cntr.UserAdd)
	user.Post("/add", cntr.UserAdd)
	user.Post("/auth", cntr.UserAuth)
	goji.Serve()
}

func init() {
	db, _ = gorm.Open("mysql", "root:@/seacher?charset=utf8&parseTime=True")
	cntr = controller.AppContext(db)
}
