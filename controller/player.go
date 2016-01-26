package controller

import (
	"encoding/json"
	"github.com/syo-sa1982/SeacherServerSide/model"
	"github.com/syo-sa1982/SeacherServerSide/util"
	"github.com/zenazn/goji/web"
	"log"
	"net/http"
	"strconv"
)

const AvoidSkillKey = 5
const LanguageSkillKey = 46

func (cntr *Controller) JobList(c web.C, w http.ResponseWriter, r *http.Request) {
	var (
		db = cntr.db
		jobSelect JobSelectAPI
	)

	var (
		Jobs = []model.JobMaster{}
		JobSkills = []model.JobSkillMaster{}
	)

	db.Order("ID", true).Find(&Jobs)
	db.Order("ID", true).Find(&JobSkills)

	jobSelect.JobMaster = Jobs
	jobSelect.JobSkillMaster = JobSkills

	log.Println(jobSelect)
	encoder := json.NewEncoder(w)
	encoder.Encode(jobSelect)

}

func (cntr *Controller) PlayerBaseMake(c web.C, w http.ResponseWriter, r *http.Request) {
	var (
		db = cntr.db
		charaMakeAPI CharaMakeAPI
		baseRolls = map[string][]int{
			"Strength"     : {6, 3},
			"Constitution" : {6, 3},
			"Power"        : {6, 3},
			"Dextality"    : {6, 3},
			"Appeal"       : {6, 3},
			"Size"         : {6, 2, 6},
			"Intelligence" : {6, 2, 6},
			"Education"    : {6, 2, 3},
		}
	)

	User := model.User{}
	r.ParseForm()
	User.UUID = r.FormValue("UUID")
	db.Find(&User)

	playerBase, history := util.CreatePlayerBase(baseRolls)
	player := util.CreatePlayerStatus(playerBase)

	log.Println(playerBase)
	log.Println(player)

	playerBase.UserID = User.ID
	playerBase.Name = User.Name
	player.UserID = User.ID
	player.JobID, _ = strconv.Atoi(r.FormValue("JobID"))

	charaMakeAPI.BaseStatus = playerBase
	charaMakeAPI.Status    = player
	charaMakeAPI.DiceHistory = history
	log.Println(charaMakeAPI)

	encoder := json.NewEncoder(w)
	encoder.Encode(charaMakeAPI)
}

func (cntr *Controller) PlayerGenerate(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = cntr.db

	User := model.User{}
	r.ParseForm()
	User.UUID = r.FormValue("UUID")
	db.Find(&User)

	log.Println(r.FormValue("data"))

	var charaMakeAPI CharaMakeAPI

	json.Unmarshal([]byte(r.FormValue("data")), &charaMakeAPI)

	baseStatus := charaMakeAPI.BaseStatus
	status     := charaMakeAPI.Status

	db.Create(&baseStatus)
	status.PlayerID = baseStatus.ID

	db.Create(&status)

	log.Println(charaMakeAPI)
}

func (cntr *Controller) SkillSetting(c web.C, w http.ResponseWriter, r *http.Request) {

	var (
	db = cntr.db
	SkillSet SkillSetAPI
)

	var (
	User = model.User{}
	PlayerBase = model.PlayerBase{}
	Player = model.PlayerStatus{}
	Job = model.JobMaster{}
	JobSkill = []model.JobSkillMaster{}
)

	r.ParseForm()
	log.Println(r.FormValue("UUID"))

	db.Model(User).Where("uuid = ?", r.FormValue("UUID")).Find(&User)
	db.Model(PlayerBase).Where("user_id = ?", User.ID).Find(&PlayerBase)
	db.Model(Player).Where("user_id = ?", User.ID).Find(&Player)
	log.Println(Player.JobID)
	db.Model(Job).Where("id = ?", Player.JobID).Find(&Job)
	db.Model(JobSkill).Where("job_id = ?", Player.JobID).Order("ID", true).Find(&JobSkill)
	log.Println(Job)

	SkillMasters := []model.SkillMaster{}
	db.Order("ID", true).Find(&SkillMasters)

	SkillMasters[AvoidSkillKey].Value = PlayerBase.Dextality * 2
	SkillMasters[LanguageSkillKey].Value = PlayerBase.Education * 5

	SkillSet.PlayerBase = PlayerBase
	SkillSet.PlayerStatus = Player
	SkillSet.SkillMaster = SkillMasters
	SkillSet.JobMaster = Job
	SkillSet.JobSkillMaster = JobSkill

	log.Println(SkillMasters)
	encoder := json.NewEncoder(w)
	encoder.Encode(SkillSet)
}

func (cntr *Controller) SkillSubmit(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = cntr.db

	r.ParseForm()
	log.Println(r.FormValue("json_api"))
	json_str := r.FormValue("json_api")
	player_id_str := r.FormValue("player_id")
	log.Println(player_id_str)
	var SkillApi []model.SkillMaster

	json.Unmarshal([]byte(json_str), &SkillApi)

	log.Println(SkillApi)

	for key := range SkillApi {
		var playerSkill model.PlayerSkill
		playerSkill.PlayerID,_ = strconv.Atoi(player_id_str)
		playerSkill.SkillID = SkillApi[key].ID
		playerSkill.Value = SkillApi[key].Value
		db.Create(&playerSkill)
	}
}
