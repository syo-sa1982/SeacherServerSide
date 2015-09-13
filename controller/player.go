package controller
import (
	"github.com/zenazn/goji/web"
	"net/http"
	"github.com/syo-sa1982/SeacherServerSide/util"
	"log"
)

func (cntr Controller) PlayerBaseMake(c web.C, w http.ResponseWriter, r *http.Request) {
	var dice = util.Roll(6, 3) // 3D6

	log.Println(dice.TotalRolls)
	log.Println(dice.TotalScore)
	log.Println(dice.RollHistory)


}