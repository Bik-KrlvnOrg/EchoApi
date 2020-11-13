package persistence

import (
	"echoApi/domain/entity"
	"echoApi/domain/repository"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Repositories struct {
	User repository.UserRepository
	db   *gorm.DB
}

func NewRepositories(dbDriver, dbUser, dbPassword, dbPort, dbHost, dbName string) (*Repositories, error) {
	connectionString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s sslmode=disable password=%s", dbHost, dbPort, dbName, dbUser, dbPassword)
	db, err := gorm.Open(dbDriver, connectionString)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	return &Repositories{
		User: NewUserRepository(db),
		db:   db,
	}, nil
}

func (repositories *Repositories) Close() error {
	return repositories.db.Close()
}

func (repositories *Repositories) AutoMigrate() error {
	return repositories.db.AutoMigrate(&entity.User{}).Error
}
