package util
import (
	"math/rand"
	"time"
)

type Dice struct {
	TotalRolls int // 振った回数
	TotalScore int // 出目の合計
	RollHistory []int // 出目の履歴
	
}

func Roll(surface int, rollcount int) Dice {
	var dice = Dice{}
	dice.TotalRolls = rollcount
	for i := 0; i < rollcount; i++ {
		rand.Seed(time.Now().UnixNano())
		dice.RollHistory = append(dice.RollHistory, rand.Intn(surface))
		dice.TotalScore += dice.RollHistory[i]
	}
	return dice
}
