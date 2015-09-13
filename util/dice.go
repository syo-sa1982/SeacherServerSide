package util
import (
	"math/rand"
	"time"
	"log"
)

type Dice struct {
	TotalScore int // 出目の合計
	RollHistory []int // 出目の履歴
	
}

func (this Dice) DiceRoll(surface int, rollcount int) (int, []int) {
	for i := 0; i < rollcount; i++ {
		rand.Seed(time.Now().UnixNano())
		this.RollHistory = append(this.RollHistory, rand.Intn(surface) + 1)
		this.TotalScore += this.RollHistory[i]
	}
	log.Println(this.RollHistory)
	return this.TotalScore, this.RollHistory
}
