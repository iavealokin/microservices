package model

//User struct ...
type User struct {
	ID       int    `json:"userid"`
	Login    string `json:"login"`
	Username string `json:"username"`
	Surname  string `json:"surname"`
	Birthday string `json:"birthday"`
	Password string `json:"password"`
}

