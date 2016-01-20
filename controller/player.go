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

	playerBase, history := createPlayerBase(baseRolls)
	player := createPlayerStatus(playerBase)

	log.Println(playerBase)
	log.Println(player)

	charaMakeAPI.BaseStatus = playerBase
	charaMakeAPI.Status    = player
	charaMakeAPI.DiceHistory = history
	log.Println(charaMakeAPI)

	encoder := json.NewEncoder(w)
	encoder.Encode(charaMakeAPI)
}

func createPlayerBase(rolls map[string][]int) (model.PlayerBase, map[string][]int) {
	var history = make(map[string][]int)
	var base model.PlayerBase

	base.Strength, history["Strength"] = util.Dice{}.DiceRoll(rolls["Strength"])
	base.Constitution, history["Constitution"] = util.Dice{}.DiceRoll(rolls["Constitution"])
	base.Power, history["Power"] = util.Dice{}.DiceRoll(rolls["Power"])
	base.Dextality, history["Dextality"] = util.Dice{}.DiceRoll(rolls["Dextality"])
	base.Appeal, history["Appeal"] = util.Dice{}.DiceRoll(rolls["Appeal"])
	base.Size, history["Size"] = util.Dice{}.DiceRoll(rolls["Size"])
	base.Intelligence, history["Intelligence"] = util.Dice{}.DiceRoll(rolls["Intelligence"])
	base.Education, history["Education"] = util.Dice{}.DiceRoll(rolls["Education"])

	return base, history
}

func (cntr *Controller) PlayerGenerate(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = cntr.db

	User := model.User{}
	r.ParseForm()
	User.UUID = r.FormValue("UUID")
	db.Find(&User)

	var BaseStatus = make(map[string]int)

	log.Println(r.Form)
	for key, value := range r.Form {
		log.Println("key:", key, " value:", value)
		if key != "UUID" || key != "JobID" {
			BaseStatus[key], _ = strconv.Atoi(value[0])
		}
	}

	var charaStatus = generatePlayerStatusMap(BaseStatus)
	log.Println(charaStatus)

	playerBase := model.PlayerBase{}
	MapToStruct(BaseStatus, &playerBase)
	playerBase.UserID = User.ID
	playerBase.Name = User.Name
	db.Create(&playerBase)

	log.Println(playerBase)

	playerStatus := model.PlayerStatus{}
	MapToStruct(charaStatus, &playerStatus)
	playerStatus.UserID = User.ID
	playerStatus.PlayerID = playerBase.ID
	playerStatus.JobID, _ = strconv.Atoi(r.FormValue("JobID"))
	playerStatus.MaxHP = charaStatus["HP"]
	playerStatus.MaxMP = charaStatus["MP"]

	db.Create(&playerStatus)

	log.Println(playerStatus)

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

func MapToStruct(m map[string]int, val interface{}) error {
	tmp, err := json.Marshal(m)
	if err != nil {
		return err
	}
	err = json.Unmarshal(tmp, val)
	if err != nil {
		return err
	}
	return nil
}


func createPlayerStatus(base model.PlayerBase) model.PlayerStatus {
	return model.PlayerStatus{
		HP              : calcDivision(2, base.Constitution, base.Size),
		MP              : base.Power,
		Sanity          : base.Power * 5,
		Luck            : base.Power * 5,
		Idea            : base.Intelligence * 5,
		Knowledge       : base.Education * 5,
		JobSkillPoint   : base.Education * 20,
		HobbySkillPoint : base.Intelligence * 10,
		DamageBonus     : base.Strength + base.Size,
	}
}

func generatePlayerStatusMap(baseStatus map[string]int) map[string]int {
	return map[string]int{
		"HP":              calcDivision(2, baseStatus["Constitution"], baseStatus["Size"]),
		"MP":              baseStatus["Power"],
		"Sanity":          baseStatus["Power"] * 5,
		"Luck":            baseStatus["Power"] * 5,
		"Idea":            baseStatus["Intelligence"] * 5,
		"Knowledge":       baseStatus["Education"] * 5,
		"JobSkillPoint":   baseStatus["Education"] * 20,
		"HobbySkillPoint": baseStatus["Intelligence"] * 10,
		"DamageBonus":     baseStatus["Strength"] + baseStatus["Size"]}
}

func calcDivision(multiple int, params ...int) int {
	var sum int = 0
	for _, param := range params {
		sum += param
	}
	if sum > 0 {
		return sum / multiple
	} else {
		return sum
	}
}
