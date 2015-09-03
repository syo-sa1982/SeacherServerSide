package main

import (
	"log"
	"github.com/zenazn/goji"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
)

var db gorm.DB

func main() {
	log.Print("main")
	user := web.New()
	goji.Handle("/user/*", user)

	user.Use(middleware.SubRouter)
	user.Get("/index", UserIndex)
	user.Get("/add", UserAdd)
	user.Post("/add", UserAdd)
	goji.Serve()
}

func init() {
	db, _ = gorm.Open("mysql", "root:@/searcher?charset=utf8&parseTime=True")
}
