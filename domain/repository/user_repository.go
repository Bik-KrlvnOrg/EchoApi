package repository

import "echoApi/domain/entity"

type UserRepository interface {
	SaveUser(*entity.User) (*entity.User, map[string]string)
	GetUsers() (*entity.Users, error)
	GetUser(uint) (*entity.User, error)
}
