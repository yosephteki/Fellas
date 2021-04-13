package Models

type UserIdea struct {
	User  User `json:"user"`
	Ideas Idea `json:"ideas"`
}
