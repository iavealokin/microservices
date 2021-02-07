package teststore

import (
	"github.com/iavealokin/microservices/MS_REST_API/app/model"
	"github.com/iavealokin/microservices/MS_REST_API/app/store"
)

//Store ..
type Store struct {
	UserRepository *UserRepository
}

func New() *Store {
return &Store{}
}


func (s *Store) User() store.UserRepository{
	if s.UserRepository!= nil{
		return s.UserRepository
	}

	s.UserRepository=&UserRepository{
		store: s,
		users: make(map[int]*model.User),
	}
		return s.UserRepository
}