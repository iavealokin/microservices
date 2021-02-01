package teststore_test

import (
	"testing"

	"github.com/iavealokin/MoneyDrive/internal/app/model"
	"github.com/iavealokin/MoneyDrive/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)


func TestUserRepository_Create(t *testing.T){
	s:= teststore.New()
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

