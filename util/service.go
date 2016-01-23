package util

import "github.com/syo-sa1982/SeacherServerSide/model"

func CreatePlayerBase(rolls map[string][]int) (model.PlayerBase, map[string][]int) {
	var history = make(map[string][]int)
	var base model.PlayerBase

	base.Strength, history["Strength"] = Dice{}.DiceRoll(rolls["Strength"])
	base.Constitution, history["Constitution"] = Dice{}.DiceRoll(rolls["Constitution"])
	base.Power, history["Power"] = Dice{}.DiceRoll(rolls["Power"])
	base.Dextality, history["Dextality"] = Dice{}.DiceRoll(rolls["Dextality"])
	base.Appeal, history["Appeal"] = Dice{}.DiceRoll(rolls["Appeal"])
	base.Size, history["Size"] = Dice{}.DiceRoll(rolls["Size"])
	base.Intelligence, history["Intelligence"] = Dice{}.DiceRoll(rolls["Intelligence"])
	base.Education, history["Education"] = Dice{}.DiceRoll(rolls["Education"])

	return base, history
}

func CreatePlayerStatus(base model.PlayerBase) model.PlayerStatus {
	return model.PlayerStatus{
		MaxHP           : calcDivision(2, base.Constitution, base.Size),
		MaxMP           : base.Power,
		HP              : calcDivision(2, base.Constitution, base.Size),
		MP              : base.Power,
		Sanity          : base.Power * 5,
		Luck            : base.Power * 5,
		Idea            : base.Intelligence * 5,
		Knowledge       : base.Education * 5,
		JobSkillPoint   : base.Education * 20,
		HobbySkillPoint : base.Intelligence * 10,
		DamageBonus     : base.Strength + base.Size,
	}
}


func calcDivision(multiple int, params ...int) int {
	var sum int = 0
	for _, param := range params {
		sum += param
	}
	if sum > 0 {
		return sum / multiple
	} else {
		return sum
	}
}