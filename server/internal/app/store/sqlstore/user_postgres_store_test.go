package store_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	model "github.com/t67y110v/software-engineering/internal/app/model/user"
	"github.com/t67y110v/software-engineering/internal/app/store"
	"github.com/t67y110v/software-engineering/internal/app/store/teststore"
)

func TestPostgresRepository_Create(t *testing.T) {
	s := teststore.NewPostgres()
	u := model.TestUser(t)
	assert.NoError(t, s.UserRepository().Create(u))
	assert.NotNil(t, u)
}

func TestPostgresRepository_FindByEmail(t *testing.T) {
	s := teststore.NewPostgres()
	email := "userTest1@test.org"
	_, err := s.UserRepository().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = email
	s.UserRepository().Create(u)
	u, err = s.UserRepository().FindByEmail(email)
	assert.NotNil(t, u)
	assert.NoError(t, err)
	assert.Equal(t, u.Email, email)

}

func TestPostgresRepository_FindByID(t *testing.T) {
	s := teststore.NewPostgres()
	ID := 1
	_, err := s.UserRepository().FindByID(strconv.Itoa(ID))
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.ID = ID
	s.UserRepository().Create(u)
	u, err = s.UserRepository().FindByID(strconv.Itoa(ID))
	assert.NotNil(t, u)
	assert.NoError(t, err)
	assert.Equal(t, u.ID, ID)
}
