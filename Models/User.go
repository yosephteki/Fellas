package Models

import (
	"Fellas/Config"
	"errors"
	"fmt"
)

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Password string `json:"password"`
}

func (u *User) TableName() string {
	return "user"
}

func Login(user *User, email string, password string) (err error) {
	if err = Config.DB.Where("email=?", email).First(user).Error; err != nil {
		return err
	}
	if user.Password != password {
		return errors.New("Wrong Password")
	}
	return nil
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
	fmt.Println(user)
	Config.DB.Save(user)
	return nil
}

func DeleteUser(user *User, id string) (err error) {
	Config.DB.Where("id=?", id).Delete(user)
	return nil
}

func GetUserIdeas(id string) (ud UserIdea, err error) {
	var user User
	var ideas []Idea
	var userIdea UserIdea
	if err = Config.DB.Where("id=?", id).First(&user).Error; err != nil {
		return ud, err
	}
	if err = Config.DB.Where("founder=?", user.Id).Find(&ideas).Error; err != nil {
		return ud, err
	}
	userIdea.User = user
	userIdea.Ideas = ideas
	return userIdea, nil

}

func GetUserIdeaJoin(id string) (ud []UserIdeaJoin, err error) {
	var userIdeaJoins []UserIdeaJoin
	rows, err := Config.DB.Table("user").
		Select("COALESCE(user.id,0) as UserId,COALESCE(user.name,'') as name,COALESCE(user.email,'') as email,COALESCE(user.phone,'') as phone,COALESCE(user.address,'') as Address,COALESCE(idea.id,0) as IdeaId,COALESCE(idea.name,'') as ideaName,COALESCE(idea.description,'') as ideaDescription,COALESCE(idea.founder,'') as founder").
		Joins("left join idea on idea.founder = user.id").Rows()
	for rows.Next() {
		var userIdeaJoin UserIdeaJoin
		err := rows.Scan(&userIdeaJoin.UserId, &userIdeaJoin.Name, &userIdeaJoin.Email, &userIdeaJoin.Phone, &userIdeaJoin.Address,
			&userIdeaJoin.IdeaId, &userIdeaJoin.IdeaName, &userIdeaJoin.IdeaDescription, &userIdeaJoin.Founder)
		if err != nil {
			fmt.Println(err)
			return ud, err
		}
		userIdeaJoins = append(userIdeaJoins, userIdeaJoin)
	}
	return userIdeaJoins, nil
}
