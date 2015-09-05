package main
import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/syo-sa1982/SeacherServerSide/model"
)


func main(){
	db, _ := gorm.Open("mysql", "root:@/seacher?charset=utf8&parseTime=True")
	db.CreateTable(&model.User{})
//	db.CreateTable(&models.Player{})
}
