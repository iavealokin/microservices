package model

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

//User struct ...
type User struct {
	ID       int    `json:"userid"`
	Login    string `json:"login"`
	Username string `json:"username"`
	Surname  string `json:"surname"`
	Birthday string `json:"birthday"`
	Password string `json:"password"`
}

var (
	db  *sql.DB
	err error
)

//InitDB func...
func InitDB() (err error) {
	db, err = sql.Open("postgres", "postgres://remote:Cfyz11005310@localhost/microservices")
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	return err

}

//GetUser func ...
func GetUser() (users []User, err error) {

	rows, err := db.Query("select id, login,name,surname,birthday from users;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		usr := new(User)
		err := rows.Scan(&usr.ID, &usr.Login, &usr.Username, &usr.Surname, &usr.Birthday)
		if err != nil {
			log.Fatal(err)
		}
		us := User{ID: usr.ID, Login: usr.Login, Username: usr.Username, Surname: usr.Surname, Birthday: usr.Birthday}
		users = append(users, us)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	for _, usr := range users {
		fmt.Printf("%s, %s, %s, %s,%s", fmt.Sprint(usr.ID), usr.Login, usr.Username, usr.Surname, usr.Birthday)
		fmt.Println("")
	}

	return users, err
}

//DropUser func..
func DropUser(userid int) (state string, errorDrop error) {
	sqlStatement := `
DELETE FROM users
WHERE id = $1;`
	if userid == 1 {
		errorDrop = errors.New("Permission denied - delete user Admin")
	} else {
		state = "ok"
		res, err := db.Exec(sqlStatement, userid)
		if err != nil {
			panic(err)
		}
		count, err := res.RowsAffected()
		if err != nil {
			panic(err)
		}
		if count == 0 {
			state = "No records delete"
			errorDrop = errors.New("No records delete")
		} else {
			state = "User is deleted"
		}
	}
	fmt.Println(errorDrop)
	return state, errorDrop
}

// AddUser func ...
func AddUser(user *User) (state string, err error) {
	sqlStatement := `
INSERT INTO users (login,name,surname,birthday,password)
values($1,$2,$3,$4,$5)`
	_, err = db.Exec(sqlStatement, user.Login, user.Username, user.Surname, user.Birthday, user.Password)
	if err != nil {
		panic(err)
	}
	return state, err
}

//UpdateUser func...
func UpdateUser(user *User) (state string, err error) {
	sqlStatement := `
UPDATE users
SET login = $1,
name = $2,
surname = $3,
birthday = $4,
password = $5
where id = $6;`
	res, err := db.Exec(sqlStatement, user.Login, user.Username, user.Surname, user.Birthday, user.Password, user.ID)
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	if count == 0 {
		state = "No records for update"
		err = errors.New("No records for update")
	} else {
		state = "User is updated"
	}
	return state, err
}
