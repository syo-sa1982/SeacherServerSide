package controller
import (
	"github.com/zenazn/goji/web"
	"net/http"
	"github.com/syo-sa1982/SeacherServerSide/util"
	"log"
	"encoding/json"
)

func (cntr Controller) PlayerBaseMake(c web.C, w http.ResponseWriter, r *http.Request) {
	var totalScore, rollHistory = util.Dice{}.DiceRoll(6, 3) // 3D6

	r.ParseForm()

	log.Print(r.Form)

	log.Println(totalScore)
	log.Println(rollHistory)

	encoder := json.NewEncoder(w)
	encoder.Encode(totalScore)
}