package sqlstore

import (
	"errors"
	"fmt"
	"log"

	"github.com/iavealokin/microservices/MS_REST_API/app/model"
)

type UserRepository struct {
	store *Store
}
//Create user... 
func (r *UserRepository) Create(u *model.User) (error) {
	return r.store.db.QueryRow(
		"INSERT INTO users (login,name,surname,birthday,password) values($1,$2,$3,$4,$5) RETURNING id",
	&u.Login,
	&u.Username,
  	&u.Surname,	
   	&u.Birthday, 
   	&u.Password,
	).Scan(&u.ID)
}

//Drop user ...
func (r *UserRepository) Drop(u *model.User) (error){
	var errorDrop error
	sqlStatement := `
DELETE FROM users
WHERE id = $1;`
	if u.ID == 1 {
		errorDrop = errors.New("Permission denied - delete user Admin")
	} else {
		res, err := r.store.db.Exec(sqlStatement, u.ID)
		if err != nil {
			panic(err)
		}
		count, err := res.RowsAffected()
		if err != nil {
			panic(err)
		}
		if count == 0 {
			errorDrop = errors.New("No records delete")
		} 
	}
	
	return errorDrop
	}

//Update user ...
func (r *UserRepository) Update(u *model.User) (error){
	var errorDrop error
	sqlStatement := `
	UPDATE users
	SET login = $1,
	name = $2,
	surname = $3,
	birthday = $4,
	password = $5
	where id = $6;`
		res, err := r.store.db.Exec(sqlStatement, u.Login, u.Username, u.Surname, u.Birthday, u.Password, u.ID)
		if err != nil {
			panic(err)
		}
		count, err := res.RowsAffected()
		if err != nil {
			panic(err)
		}
		if count == 0 {
			errorDrop = errors.New("No records for update")
		} 
		return errorDrop
	}
	//Get users list
	func (r *UserRepository) Get() ([] model.User, error){
		
		sqlStatement := `select id,login,name,surname,birthday from users;`
		var users [] model.User
		rows,err := r.store.db.Query(sqlStatement)
			if err != nil {
				log.Fatal(err)
			}
			defer rows.Close()
			for rows.Next() {
				usr := new(model.User)
				err := rows.Scan(&usr.ID, &usr.Login, &usr.Username, &usr.Surname, &usr.Birthday)
				if err != nil {
					log.Fatal(err)
				}
				us := model.User{ID: usr.ID, Login: usr.Login, Username: usr.Username, Surname: usr.Surname, Birthday: usr.Birthday}
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



