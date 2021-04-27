package main

import (
	"Fellas/Config"
	"Fellas/Routes"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	var err error
	fmt.Println("HELLO WORLD!")

	Config.DB, err = gorm.Open("mysql", Config.DBURL(Config.BuildConfig()))
	if err != nil {
		fmt.Println("Status : ", err)
	}
	// Config.DB.AutoMigrate(&Models.Idea{})
	// Config.DB.AutoMigrate(&Models.User{})

	defer Config.DB.Close()
	r := Routes.SetupRouter()
	r.Run()
}
