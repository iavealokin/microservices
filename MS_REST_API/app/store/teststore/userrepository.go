package teststore

import (
	"github.com/iavealokin/microservices/MS_REST_API/app/model"
)


type UserRepository struct {
	store *Store
	users map[int]*model.User
}

//Create test user
func (r *UserRepository) Create(u *model.User) error{
	u.ID = len(r.users) + 1
	r.users[u.ID]= u
	
	return nil
}
