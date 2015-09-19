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

	r.ParseForm()

	var totalScores map[string]int = make(map[string]int)
	var historys map[string][]int = make(map[string][]int)

	for key, value := range r.Form {
		log.Println("key:", key, " value:", value)
		totalScores[key] ,historys[key] = parameterGenerate(r, key)
	}

	log.Println(totalScores)
	log.Println(historys)

	encoder := json.NewEncoder(w)
	encoder.Encode(totalScores)
}

func (cntr Controller) PlayerGenerate(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = cntr.db
	User := model.User{}
	r.ParseForm()
	User.UUID = r.FormValue("uuid")
	db.Find(&User)

	Player := model.PlayerBase{UserID:User.ID,Name:User.Name}

	log.Println(Player)

}


func parameterGenerate(r *http.Request, key string) (int, []int) {

	split := strings.Split(r.FormValue(key), ",")
	serface, _ := strconv.Atoi(split[0])
	rollCount, _ := strconv.Atoi(split[1])

	return util.Dice{}.DiceRoll(serface, rollCount)
}