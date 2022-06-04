package main

import (
	"gin-project/docs"
	"gin-project/initialize"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	docs.SwaggerInfo.Title = "bbs"
	docs.SwaggerInfo.Description = "This is bbs's Golang backend."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	err := initialize.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer initialize.Close()

	r := initialize.SetupRouter()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8889")
}
