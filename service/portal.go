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
