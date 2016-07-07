package logic

import (
	"github.com/jinzhu/gorm"
	"jasonghw/xproject/DemoService/model"
)

func GetUidGroupCount(db *gorm.DB, uid string) (int, error) {
	count := 0
	err := db.Model(model.SportsGroup{}).Where("user_id = ?", uid).Count(&count).Error
	return count, err
}

type CountRet struct {
	Count int
	Err   error
}
