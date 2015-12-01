package controller
import (
	"github.com/zenazn/goji/web"
	"net/http"
	"github.com/syo-sa1982/SeacherServerSide/model"
	"log"
	"html/template"
)



func (cntr Controller) MainIndex(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = cntr.db
	Users := []model.User{}
	db.Find(&Users)
	log.Print("index")
	tpl = template.Must(template.ParseFiles("view/user/index.html"))
	tpl.Execute(w, Users)
}