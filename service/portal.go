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
func CreateNotification(notification *model.Notification) (err error) {
	err = global.DB.Create(notification).Error
	return err
}
func GetLikeNotification(userID uint64, off uint64, lim uint64) (
	notifications []model.Notification, count uint64) {
	global.DB.Order("create_at desc").Where("user_id = ? and type < 2", userID).
		Limit(lim).Offset(off).Find(&notifications).
		Limit(-1).Offset(-1).Count(&count)
	return notifications, count
}
func GetCommentNotification(userID uint64, off uint64, lim uint64) (
	notifications []model.Notification, count uint64) {
	global.DB.Order("create_at desc").Where("user_id = ? and type = 2", userID).
		Limit(lim).Offset(off).Find(&notifications).
		Limit(-1).Offset(-1).Count(&count)
	return notifications, count
}
func CreateSysMessage(sm *model.SysMessage) (err error) {
	err = global.DB.Create(sm).Error
	return err
}
func GetSysMessageCount(userID uint64, t int, date string) (
	count int) {
	global.DB.Model(&model.SysMessage{}).Where("user_id = ? and type = ?", userID, t).
		Where("date = ?", date).Count(&count)
	return count
}
func GetSysMessages(userID uint64, off uint64, lim uint64) (
	sysMessages []model.SysMessage, count uint64) {
	global.DB.Order("create_at desc").Where("user_id = ?", userID).
		Limit(lim).Offset(off).Find(&sysMessages).
		Limit(-1).Offset(-1).Count(&count)
	return sysMessages, count
}
