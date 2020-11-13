package app

import (
	"echoApi/domain/entity"
	"echoApi/domain/repository"
)

type EchoAppInterface interface {
	SaveEchos(*entity.Item) (*entity.Item, error)
	GetEchos() (*entity.Items, error)
}

type echoApp struct {
	repository repository.EchoRepository
}

var _ EchoAppInterface = &echoApp{}

func (e echoApp) SaveEchos(item *entity.Item) (*entity.Item, error) {
	return e.repository.SaveEchos(item)
}

func (e echoApp) GetEchos() (*entity.Items, error) {
	return e.repository.GetEchos()
}
