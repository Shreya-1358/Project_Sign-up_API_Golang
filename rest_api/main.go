package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	Config "repos/project/config"
	"repos/project/model"
	"repos/project/routes"
)

var err error

func init() {
	Config.DB, err = gorm.Open(mysql.Open(Config.DbURL()), &gorm.Config{})
	if err != nil {
		log.Fatal("Status:", err)
	}

	Config.DB.AutoMigrate(&model.User{})

}

func main() {

	r := routes.SetupRouter()
	r.Run()
}
