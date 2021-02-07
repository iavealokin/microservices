package teststore_test

import (
	"testing"

	"github.com/iavealokin/microservices/MS_REST_API/app/model"
	"github.com/iavealokin/microservices/MS_REST_API/app/store/teststore"
	"github.com/stretchr/testify/assert"
)


func TestUserRepository_Create(t *testing.T){
	s:= teststore.New()
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_Drop(t *testing.T){
	s:= teststore.New()
	u:=model.TestUser(t)
	assert.NoError(t,s.User().Drop(u))
	assert.NotNil(t,u)
}
