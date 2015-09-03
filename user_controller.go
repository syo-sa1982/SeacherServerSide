package main

import (
	"github.com/zenazn/goji/web"
	"net/http"
	"html/template"
	"log"
	"github.com/syo-sa1982/SeacherServerSide/models"
	"strconv"
)

var tpl *template.Template
const Password = "user:user"

type FormData struct{
	User models.User
	Mess string
}

func UserIndex(c web.C, w http.ResponseWriter, r *http.Request){
	Users := [] models.User{}
	db.Find(&Users)
	log.Print("index")
	tpl = template.Must(template.ParseFiles("views/user/index.html"))
	tpl.Execute(w, Users)
}

func UserAdd(c web.C, w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	log.Print("add")
	log.Print(r.Form)
	User := models.User{UUID: r.FormValue("uuid")}
	User.RollCount, _ = strconv.Atoi(r.FormValue("roll_count"))
	db.Create(&User)
}