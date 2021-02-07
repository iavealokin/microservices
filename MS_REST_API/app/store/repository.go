package store

import "github.com/iavealokin/microservices/MS_REST_API/app/model"

type UserRepository interface {
	Create(*model.User) error
	Drop(*model.User) error
	Update(*model.User) error
	Get() ([] model.User, error)
	UserLogin(string, string)  error
	/*FindByEmail(string) (*model.User, error)
	FindById(int) (*model.User, error)*/
}