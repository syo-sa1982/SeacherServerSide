package main

import (
	"log"
	"github.com/zenazn/goji"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

var db gorm.DB

func main() {
	log.Print("main")
	goji.Serve()
}

func init() {
	db, _ = gorm.Open("mysql", "root:@/searcher?charset=utf8&parseTime=True")
}
