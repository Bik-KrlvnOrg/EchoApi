package persistence

import (
	"echoApi/domain/entity"
	"echoApi/domain/repository"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"strings"
)

type UserRepo struct {
	db *gorm.DB
}

var _ repository.UserRepository = &UserRepo{}

func NewUserRepository(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}

func (repo UserRepo) SaveUser(user *entity.User) (*entity.User, map[string]string) {
	errors := map[string]string{}
	err := repo.db.Debug().Create(&user).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			errors["username"] = "already exits"
		}
		errors["db_errors"] = "db error"
		return nil, errors
	}
	return user, nil
}

func (repo UserRepo) GetUsers() (*entity.Users, error) {
	var users entity.Users
	err := repo.db.Debug().Find(&users).Error
	if err != nil {
		return nil, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("no records available")
	}

	return &users, nil
}

func (repo UserRepo) GetUser(id uint) (*entity.User, error) {
	var user entity.User
	err := repo.db.Debug().Where("id = ?", id).Take(&user).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("user not found")
	}
	return &user, nil
}
