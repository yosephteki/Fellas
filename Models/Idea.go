package Models

import "Fellas/Config"

type Idea struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Founder     string `json:"founder"`
	// Contributor
}

func (u *Idea) TableName() string {
	return "idea"
}

func GetAllIdeas(idea *[]Idea) (err error) {
	if err = Config.DB.Find(idea).Error; err != nil {
		return err
	}
	return nil
}

func GetIdeaByUserId(idea *Idea, userId string) (err error) {
	if err = Config.DB.Where("founder=?", userId).First(idea).Error; err != nil {
		return err
	}
	return nil

}
