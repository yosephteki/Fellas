package main

import (
	"Fellas/Config"
	"Fellas/Routes"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	var err error
	fmt.Println("HELLO WORLD!")

	Config.DB, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres sslmode=disable password=admin")
	if err != nil {
		fmt.Println("Status : ", err)
	}
	// Config.DB.AutoMigrate(&Models.Idea{})
	// Config.DB.AutoMigrate(&Models.User{})

	defer Config.DB.Close()
	r := Routes.SetupRouter()
	r.Run()
}
