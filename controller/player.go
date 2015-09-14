package controller
import (
	"github.com/zenazn/goji/web"
	"net/http"
	"github.com/syo-sa1982/SeacherServerSide/util"
	"log"
	"encoding/json"
	"strings"
	"strconv"
)

func (cntr Controller) PlayerBaseMake(c web.C, w http.ResponseWriter, r *http.Request) {
	r.ParseForm()


	for key, value := range r.Form {
		log.Println("key:", key, " value:", value)
	}

	var totalScore, rollHistory = parameterGenerate(r, "Strength")


	log.Println(totalScore)
	log.Println(rollHistory)

	encoder := json.NewEncoder(w)
	encoder.Encode(totalScore)
}



func parameterGenerate(r *http.Request, key string) (int, []int) {

	split := strings.Split(r.FormValue(key), ",")
	serface, _ := strconv.Atoi(split[0])
	rollCount, _ := strconv.Atoi(split[1])

	return util.Dice{}.DiceRoll(serface, rollCount)
}