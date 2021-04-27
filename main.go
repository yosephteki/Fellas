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

	Config.DB, err = gorm.Open("postgres", "host=ec2-54-166-167-192.compute-1.amazonaws.com port=5432 user=nrojvokwxfufac dbname=d1ihg08luvjgbb sslmode=disable password=25bdb986f6d7469052ab3468b88b900886a70a510aaf19d5b580b65bba36b624")
	if err != nil {
		fmt.Println("Status : ", err)
	}
	// Config.DB.AutoMigrate(&Models.Idea{})
	// Config.DB.AutoMigrate(&Models.User{})

	defer Config.DB.Close()
	r := Routes.SetupRouter()
	r.Run()
}
