package util

import (
	"log"
	"math/rand"
	"time"
)

type Dice struct {
	TotalScore  int   // 出目の合計
	RollHistory []int // 出目の履歴

}

func (this Dice) DiceRoll(rollData []int) (int, []int) {
	surface   := rollData[0]
	rollcount := rollData[1]

	option := 0
	if(len(rollData) > 2) {
		option = rollData[2]
	}

	for i := 0; i < rollcount; i++ {
		rand.Seed(time.Now().UnixNano())
		this.RollHistory = append(this.RollHistory, rand.Intn(surface)+1)
		this.TotalScore += this.RollHistory[i]
	}
	this.TotalScore += option
	log.Println(this.RollHistory)
	return this.TotalScore, this.RollHistory
}
