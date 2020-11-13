package app

import (
	"echoApi/domain/entity"
	"echoApi/domain/repository"
)

type UserAppInterface interface {
	SaveUser(*entity.User) (*entity.User, map[string]string)
	GetUsers() (*entity.Users, error)
	GetUser(uint) (*entity.User, error)
}

type userApp struct {
	repository repository.UserRepository
}

var _ UserAppInterface = &userApp{}

func (e *userApp) SaveUser(user *entity.User) (*entity.User, map[string]string) {
	return e.repository.SaveUser(user)
}

func (e *userApp) GetUsers() (*entity.Users, error) {
	return e.repository.GetUsers()
}

func (e *userApp) GetUser(id uint) (*entity.User, error) {
	return e.repository.GetUser(id)
}
