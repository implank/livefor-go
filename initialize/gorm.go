package initialize

import (
	"fmt"
	"gin-project/global"
	"gin-project/model"

	"github.com/jinzhu/gorm"
)

func InitMySQL() (err error) {
	addr, port, username, password, database := "127.0.0.1", "3306", "root", "123456", "se"
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, addr, port, database)
	global.DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	global.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		// user
		&model.User{},

		// Post
		&model.Post{},
		&model.PostLike{},
		&model.Comment{},
		&model.CommentLike{},
		&model.Tag{},
		&model.PostTag{},

		//portal
		&model.Greenbird{},
	)
	return global.DB.DB().Ping()
}
func Close() {
	global.DB.Close()
}
