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

//Drop test user..
func (r *UserRepository) Drop(u *model.User) error{
	u.ID = len(r.users) + 1
	r.users[u.ID]= u
	
	return nil
}
//Get users list
func (r *UserRepository) Get() ([]model.User,error){
	
	
	return nil,nil
}
//UserLogin ...
func (r *UserRepository) UserLogin(string,string) (error){
	return nil
}
// Update ...
func (r *UserRepository) Update(u *model.User) error{
	u.ID = len(r.users) + 1
	r.users[u.ID]= u
	
	return nil
}
