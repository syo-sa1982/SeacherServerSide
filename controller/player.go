package controller
import (
	"github.com/zenazn/goji/web"
	"net/http"
	"github.com/syo-sa1982/SeacherServerSide/util"
	"log"
	"encoding/json"
	"strings"
	"strconv"
	"github.com/syo-sa1982/SeacherServerSide/model"
	"fmt"
)

func (cntr Controller) PlayerBaseMake(c web.C, w http.ResponseWriter, r *http.Request) {

	charaMakeAPI := cntr.charaStatus
	r.ParseForm()

	var totalScores = make(map[string]int)
	var history = make(map[string][]int)

	for key, value := range r.Form {
		log.Println("key:", key, " value:", value)
		totalScores[key] , history[key] = generateBaseStatus(r, key)
	}
	var charaStatus = generatePlayerStatusMap(totalScores)

	charaMakeAPI.BaseStatus = totalScores
	charaMakeAPI.CharaStatus = charaStatus
	charaMakeAPI.DiceHistory = history
	log.Println(charaMakeAPI)

	encoder := json.NewEncoder(w)
	encoder.Encode(charaMakeAPI)
}

func (cntr Controller) PlayerGenerate(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = cntr.db
	User := model.User{}
	r.ParseForm()
	User.UUID = r.FormValue("UUID")
	db.Find(&User)

	var BaseStatus = make(map[string]int)
	log.Println(r.Form)
	for key, value := range r.Form {
		log.Println("key:", key, " value:", value)
		if key != "UUID" {
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
	playerStatus.MaxHP = charaStatus["HP"]
	playerStatus.MaxMP = charaStatus["MP"]

	db.Create(&playerStatus)

	log.Println(playerStatus)

}

func (cntr Controller) SkillSetting(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = cntr.db
	var SkillSet = cntr.SkillSetAPI

	User := model.User{}
	Player := model.PlayerStatus{}
	r.ParseForm()
	log.Println(r.FormValue("UUID"))

	db.Model(User).Where("uuid = ?", r.FormValue("UUID")).Find(&User)
	log.Println(User.ID)
	db.Model(Player).Where("user_id = ?", User.ID).Find(&Player)

	SkillSet.PlayerStatus = Player

	skillMaster := []model.SkillMaster{}
	db.Find(&skillMaster)

	SkillSet.SkillMaster = getSkillMasterMap(skillMaster)

	log.Println(SkillSet)
	encoder := json.NewEncoder(w)
	encoder.Encode(SkillSet)
}

func (cntr Controller) PlayerList(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = cntr.db
	User := model.User{}
	Player := model.PlayerStatus{}
	r.ParseForm()
	log.Println(r.FormValue("UUID"))

	db.Model(User).Where("uuid = ?", r.FormValue("UUID")).Find(&User)
	log.Println(User.ID)
	db.Model(Player).Where("user_id = ?", User.ID).Find(&Player)

	log.Println(Player)

	encoder := json.NewEncoder(w)
	encoder.Encode(Player)
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



func generateBaseStatus(r *http.Request, key string) (int, []int) {
	split := strings.Split(r.FormValue(key), ",")
	serface, _ := strconv.Atoi(split[0])
	rollCount, _ := strconv.Atoi(split[1])
	return util.Dice{}.DiceRoll(serface, rollCount)
}

func generatePlayerStatusMap(baseStatus map[string]int) map[string]int {
	return map[string]int{
		"HP" : calcDivision(2, baseStatus["Constitution"], baseStatus["Size"]),
		"MP" : baseStatus["Power"],
		"Sanity" : baseStatus["Power"] * 5,
		"Luck" : baseStatus["Power"] * 5,
		"Idea" : baseStatus["Intelligence"] * 5,
		"Knowledge" : baseStatus["Education"] * 5,
		"JobSkillPoint" : baseStatus["Education"] * 20,
		"HobbySkillPoint" : baseStatus["Intelligence"] * 10,
		"DamageBonus" : baseStatus["Strength"] + baseStatus["Size"]}
}

func calcDivision(multiple int, params ...int) int {
	var sum int = 0;
	for _ , param := range params { sum += param }
	if sum > 0 { return sum / multiple
	} else { return sum }
}

func getSkillMasterMap(skillMasters []model.SkillMaster) map[string]model.SkillMaster {
	skillmap := map[string]model.SkillMaster{}

	for _,skillMaster := range skillMasters {
		skillmap[fmt.Sprint(skillMaster.ID)] = skillMaster
	}

	return skillmap
}