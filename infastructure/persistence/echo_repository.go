package persistence

import (
	"echoApi/domain/entity"
	"echoApi/domain/repository"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type EchoRepo struct {
	db *gorm.DB
}

var _ repository.EchoRepository = &EchoRepo{}

func NewEchoRepository(db *gorm.DB) *EchoRepo {
	return &EchoRepo{db}
}

func (repo EchoRepo) SaveEchos(item *entity.Item) (*entity.Item, error) {
	err := repo.db.Create(&item).Error
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (repo EchoRepo) GetEchos() (*entity.Items, error) {
	var items entity.Items
	err := repo.db.Find(&items).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("data is empty")
	}
	return &items, nil
}
