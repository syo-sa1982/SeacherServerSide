package controller
import (
	"github.com/jinzhu/gorm"
)

type Controller struct {
	db gorm.DB
}

func AppContext(db gorm.DB) Controller {
	return Controller{db: db}
}

