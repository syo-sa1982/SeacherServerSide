package main
import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/syo-sa1982/SeacherServerSide/model"
)


func main(){
	db, _ := gorm.Open("mysql", "root:@/seacher?charset=utf8&parseTime=True")
//	db.DropTable(&model.User{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.User{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.PlayerBase{})
}
