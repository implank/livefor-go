package service

import (
	"gin-project/global"
	"gin-project/model"
)

func SaveGreenbird(greenBirds []model.Greenbird) (err error) {
	global.DB.Delete(&model.Greenbird{})
	for _, greenBird := range greenBirds {
		err = global.DB.Create(&greenBird).Error
		if err != nil {
			return err
		}
	}
	return
}
func GetGreenbirds() (greenbirds []model.Greenbird, err error) {
	err = global.DB.Order("order").Find(&greenbirds).Error
	return greenbirds, err
}
func CreateNotification(*model.Notification) (err error) {
	err = global.DB.Create(&model.Notification{}).Error
	return err
}
func GetNotification(userID uint64, off uint64, lim uint64) (
	notifications []model.Notification, count uint64) {
	global.DB.Order("CreateAt desc").Where("user_id = ?", userID).
		Limit(lim).Offset(off).Find(&notifications).
		Limit(-1).Offset(-1).Count(&count)
	return notifications, count
}
