package controller
import (
	"net/http"
	"github.com/zenazn/goji/web"
	"github.com/syo-sa1982/SeacherServerSide/model"
	"log"
	"html/template"
	"strconv"
	"encoding/json"
)

var tpl *template.Template
const Password = "user:user"

type FormData struct{
	User model.User
	Mess string
}

func (cntr Controller) UserIndex(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = cntr.db
	Users := [] model.User{}
	db.Find(&Users)
	log.Print("index")
	tpl = template.Must(template.ParseFiles("view/user/index.html"))
	tpl.Execute(w, Users)
}

func (cntr Controller) UserAdd(c web.C, w http.ResponseWriter, r *http.Request){
	var db = cntr.db
	r.ParseForm()
	log.Print("add")

	log.Print(w.Header())
	log.Print(r.Form)
	User := model.User{UUID: r.FormValue("uuid")}
	User.RollCount, _ = strconv.Atoi(r.FormValue("roll_count"))
	db.Create(&User)

	encoder := json.NewEncoder(w)
	encoder.Encode(User)
}

func (cntr Controller) UserAuth(c web.C, w http.ResponseWriter, r *http.Request){
	var db = cntr.db
	log.Print("Auth")
	User := model.User{}
	r.ParseForm()
	User.UUID = r.FormValue("uuid")
	db.Find(&User)

	encoder := json.NewEncoder(w)
	encoder.Encode(User)
}