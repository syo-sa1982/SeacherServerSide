package controller
import (
	"github.com/zenazn/goji/web"
	"net/http"
	"github.com/syo-sa1982/SeacherServerSide/model"
	"log"
	"encoding/json"
)



func (cntr Controller) UserInfo(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = cntr.db
	var UserInfo UserInfoAPI
	var (
		User = model.User{}
		PlayerBase = model.PlayerBase{}
		Player = model.PlayerStatus{}
		Job = model.JobMaster{}
	)

	r.ParseForm()
	log.Println(r.FormValue("UUID"))

	db.Model(User).Where("uuid = ?", r.FormValue("UUID")).Find(&User)
	db.Model(PlayerBase).Where("user_id = ?", User.ID).Find(&PlayerBase)
	db.Model(Player).Where("user_id = ?", User.ID).Find(&Player)

	log.Println(Player.JobID)
	db.Model(Job).Where("id = ?", Player.JobID).Find(&Job)

	UserInfo.User = User
	UserInfo.Job = Job
	UserInfo.PlayerBase = PlayerBase
	UserInfo.PlayerStatus = Player

	encoder := json.NewEncoder(w)
	encoder.Encode(UserInfo)

}

func (cntr *Controller) PlayerList(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = cntr.db
	user := model.User{}
	job := model.JobMaster{}
	playerBase := model.PlayerBase{}
	playerStatus := model.PlayerStatus{}
	r.ParseForm()
	log.Println(r.FormValue("UUID"))

	db.Model(user).Where("uuid = ?", r.FormValue("UUID")).Find(&user)
	db.Model(playerStatus).Where("user_id = ?", user.ID).Find(&playerStatus)
	db.Model(playerBase).Where("user_id = ?", user.ID).Find(&playerBase)
	db.Model(job).Where("job_id = ?", playerStatus.JobID).Find(&job)

	log.Println(playerStatus)

	encoder := json.NewEncoder(w)
	encoder.Encode(playerStatus)
}

func (cntr Controller) ScenarioList(c web.C, w http.ResponseWriter, r *http.Request) {

}