package service

import (
	"gin-project/global"
	"gin-project/model"
)

func QueryUserByUsername(username string) (user model.User, notFound bool) {
	notFound = global.DB.Where("username = ?", username).First(&user).RecordNotFound()
	return user, notFound
}
func QueryUserByEmail(email string) (user model.User, notFound bool) {
	notFound = global.DB.Where("username = ?", email).First(&user).RecordNotFound()
	return user, notFound
}
func QueryUserByUserID(userID uint64) (user model.User, notFound bool) {
	notFound = global.DB.Where("user_id = ?", userID).First(&user).RecordNotFound()
	return user, notFound
}
func UpdateUserPassword(user *model.User, newPassword string) error {
	user.Password = newPassword
	err := global.DB.Save(user).Error
	return err
}
func UpdateUser(user *model.User, username string, password string, email string, info string) error {
	user.Username = username
	user.Password = password
	user.Email = email
	err := global.DB.Save(user).Error
	return err
}
func CreateUser(user *model.User) (err error) {
	if err = global.DB.Create(&user).Error; err != nil {
		return err
	}
	return
}
