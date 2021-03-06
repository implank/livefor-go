package service

import (
	"gin-project/global"
	"gin-project/model"
	"gin-project/utils"
	"time"
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
func UpdateUser(user *model.User) error {
	err := global.DB.Save(user).Error
	return err
}
func CreateUser(user *model.User) (err error) {
	if err = global.DB.Create(&user).Error; err != nil {
		return err
	}
	return
}
func GetBannedUsers() (users []model.User, err error) {
	err = global.DB.Where("bandate > ?", time.Now()).Find(&users).Error
	return users, err
}

var MAXLEVEL = 10
var EXPGAP = []int{0, 30, 60, 150, 450, 1050, 2100, 3600, 5500, 10500}

func UpdateUserExp(userID uint64, exp int) {
	user, _ := QueryUserByUserID(userID)
	user.Exp += exp
	if exp >= 0 {
		for i := user.Level; i < MAXLEVEL; i++ {
			if user.Exp >= EXPGAP[i] {
				user.Level = i + 1
				user.Exp -= EXPGAP[i]
				count := GetSysMessageCount(userID, 3, "")
				if count == 0 {
					sm := model.SysMessage{
						UserID: userID,
						Date:   time.Now().Format(utils.DAYFORMAT),
						Type:   3,
						Times:  i + 1,
					}
					CreateSysMessage(&sm)
				}
			} else {
				break
			}
		}
	} else {
		for i := user.Level; i > 0; i-- {
			if user.Exp < 0 {
				user.Level = i - 1
				user.Exp += EXPGAP[i-1]
			} else {
				break
			}
		}
	}
	UpdateUser(&user)
}
