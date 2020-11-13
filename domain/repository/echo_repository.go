package repository

import "echoApi/domain/entity"

type EchoRepository interface {
	SaveEchos(*entity.Item) (*entity.Item, error)
	GetEchos() (*entity.Items, error)
}
