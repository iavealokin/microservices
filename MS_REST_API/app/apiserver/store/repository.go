package store

import "github.com/iavealokin/microservices/MS_REST_API/app/model"

type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
	FindById(int) (*model.User, error)
}