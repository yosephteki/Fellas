package Models

type UserIdea struct {
	User  User   `json:"user"`
	Ideas []Idea `json:"ideas"`
}

type UserIdeaJoin struct {
	UserId          string `json:"UserId"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	Address         string `json:"Address"`
	IdeaId          string `json:"IdeaId"`
	IdeaName        string `json:"ideaName"`
	IdeaDescription string `json:"ideaDescription"`
	Founder         string `json:"founder"`
}
