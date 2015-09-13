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

	log.Print(r.Form)
	r.FormValue("Strength")
	split := strings.Split(r.FormValue("Strength"), ",")

	serface, _ := strconv.Atoi(split[0])
	rollCount, _ := strconv.Atoi(split[1])

	var totalScore, rollHistory = util.Dice{}.DiceRoll(serface, rollCount) // 3D6

	log.Print(split)

	log.Println(totalScore)
	log.Println(rollHistory)

	encoder := json.NewEncoder(w)
	encoder.Encode(totalScore)
}