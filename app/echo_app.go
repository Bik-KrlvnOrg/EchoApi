package app

import (
	"echoApi/domain/entity"
	"echoApi/domain/repository"
)

type EchoAppInterface interface {
	SaveUser(*entity.User) (*entity.User, map[string]string)
	GetUsers() (*entity.Users, error)
	GetUser(uint) (*entity.User, error)
}

type echoApp struct {
	repository repository.UserRepository
}

var _ EchoAppInterface = &echoApp{}

func (e *echoApp) SaveUser(user *entity.User) (*entity.User, map[string]string) {
	return e.repository.SaveUser(user)
}

func (e *echoApp) GetUsers() (*entity.Users, error) {
	return e.repository.GetUsers()
}

func (e *echoApp) GetUser(id uint) (*entity.User, error) {
	return e.repository.GetUser(id)
}
