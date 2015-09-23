package controller
import (
	"github.com/jinzhu/gorm"
)

type Controller struct {
	db gorm.DB
	charaStatus CharaMakeAPI
}

type CharaMakeAPI struct {
	BaseStatus, CharaStatus map[string]int
	DiceHistory map[string][]int
}

func AppContext(db gorm.DB) Controller {
	return Controller{db: db}
}

