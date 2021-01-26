package model

//User struct ...
type User struct {
	Id      int
	Name    string
	Surname string
	Vip     int
}

//GetUser func ...
func GetUser() (users []User, err error) {
	users = []User{
		{1, "Александр", "Николаев", 0},
	}
	return
}
