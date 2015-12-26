package admin
import (
	"github.com/jinzhu/gorm"
	"github.com/zenazn/goji/web"
	"net/http"
	"github.com/syo-sa1982/SeacherServerSide/model"
	"log"
	"html/template"
)

var tpl *template.Template

type Controller struct {
	db          gorm.DB
}

func AppContext(db gorm.DB) Controller {
	return Controller{db: db}
}


func (cntr *Controller) AdminIndex(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = cntr.db
	Users := []model.User{}
	db.Find(&Users)
	log.Print("index")
	tpl = template.Must(template.ParseFiles("view/admin/index.html"))
	tpl.Execute(w, Users)
}