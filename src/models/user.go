package models

type UserForm struct {
	Name     string `json:"username"`
	Document string `json:"document"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
