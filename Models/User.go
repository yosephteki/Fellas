package Models

import (
	"Fellas/Config"
	"fmt"
)

type User struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

func (u *User) TableName() string {
	return "user"
}

func GetAllUsers(user *[]User) (err error) {
	if err = Config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserById(user *User, id string) (err error) {
	if err = Config.DB.Where("id=?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

func CreateUser(user *User) (err error) {
	if err = Config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *User, id string) (err error) {
	Config.DB.Save(user)
	return nil
}

func DeleteUser(user *User, id string) (err error) {
	Config.DB.Where("id=?", id).Delete(user)
	return nil
}

func GetUserIdeas(id string) (ud UserIdea, err error) {
	var user User
	var ideas Idea
	var userIdea UserIdea
	if err = Config.DB.Where("id=?", id).First(&user).Error; err != nil {
		return ud, err
	}
	if err = Config.DB.Where("founder=?", user.Id).Find(&ideas).Error; err != nil {
		return ud, err
	}
	fmt.Println(user)
	userIdea.User = user
	userIdea.Ideas = ideas
	return userIdea, nil

}
