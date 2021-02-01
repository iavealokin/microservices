package sqlstore

import (
	"database/sql"

	"github.com/iavealokin/microservices/MS_REST_API/app/model"
	"github.com/iavealokin/microservices/MS_REST_API/app/store"
)

type UserRepository struct {
	store *Store
}
//Create user... 
func (r *UserRepository) Create(u *model.User) (error) {
	return r.store.db.QueryRow(
		"INSERT INTO users (login,name,surname,birthday,password)
values($1,$2,$3,$4,$5) RETURNING id",
	&u.Login,
	&u.Username,
  	&u.Surname,	
   	&u.Birthday, 
   	&u.Password,
	).Scan(&u.ID)
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE email=$1",
		email,
		).Scan(
			&u.ID,
			&u.Email,
			&u.EncryptedPassword,
	); err != nil{
		if err == sql.ErrNoRows{
			return nil, store.ErrRecordNotFound
		}
		return nil,err
	}
	return u, nil
}

func (r *UserRepository) FindById(id int) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE id=$1",
		id,
		).Scan(
			&u.ID,
			&u.Email,
			&u.EncryptedPassword,
	); err != nil{
		if err == sql.ErrNoRows{
			return nil, store.ErrRecordNotFound
		}
		return nil,err
	}
	return u, nil
}