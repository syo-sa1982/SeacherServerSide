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
)

func (cntr Controller) PlayerBaseMake(c web.C, w http.ResponseWriter, r *http.Request) {

	charaMakeAPI := cntr.charaStatus
	r.ParseForm()

	var totalScores map[string]int = make(map[string]int)
	var history map[string][]int = make(map[string][]int)

	for key, value := range r.Form {
		log.Println("key:", key, " value:", value)
		totalScores[key] , history[key] = generateBaseStatus(r, key)
	}
	var charaStatus = generateCharaStatus(totalScores)

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
	User.UUID = r.FormValue("uuid")

	log.Println(r.Form)

//	var charaStatus = generateCharaStatus(r.Form)

//	log.Println(charaStatus)
	db.Find(&User)

	Player := model.PlayerBase{UserID:User.ID,Name:User.Name}

	log.Println(Player)

}


func generateBaseStatus(r *http.Request, key string) (int, []int) {
	split := strings.Split(r.FormValue(key), ",")
	serface, _ := strconv.Atoi(split[0])
	rollCount, _ := strconv.Atoi(split[1])
	return util.Dice{}.DiceRoll(serface, rollCount)
}

func generateCharaStatus(baseStatus map[string]int) map[string]int {
	return map[string]int{
		"HP" : calcDivision(2, baseStatus["Constitution"], baseStatus["Size"]),
		"MP" : baseStatus["Power"],
		"Sanity" : baseStatus["Power"] * 5,
		"Luck" : baseStatus["Power"] * 5,
		"Idea" : baseStatus["Intelligence"] * 5,
		"Knowledge" : baseStatus["Education"] * 5,
		"JopSkillPoint" : baseStatus["Education"] * 20,
		"HobbySkillPoint" : baseStatus["Intelligence"] * 10,
		"DamageBonus" : baseStatus["Strength"] + baseStatus["Size"]}
}

func calcDivision(multiple int, params ...int) int {
	var sum int = 0;
	for _ , param := range params { sum += param }
	if sum > 0 { return sum / multiple
	} else { return sum }
}
