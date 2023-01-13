package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	Config "repos/project/config"
	"repos/project/model"
	"repos/project/routes"
)

var err error

func main() {
	fmt.Println("start")
	fmt.Println("db:", Config.DbURL(Config.BuildDBConfig()))
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	fmt.Println("db end")
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&model.User{})
	r := routes.SetupRouter()
	//running
	r.Run()
}
