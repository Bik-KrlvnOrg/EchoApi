package interfaces

import (
	"echoApi/app"
	"echoApi/domain/entity"
	"echoApi/dto"
	"log"
	"net/http"
)

type User struct {
	log *log.Logger
	app app.UserAppInterface
}

func NewUser(log *log.Logger, appInterface app.UserAppInterface) *User {
	return &User{log, appInterface}
}

func (echo *User) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		user := createUser(res, req)
		data, err := echo.CreateUser(user)
		if err != nil {
			echo.log.Println(err)
			http.Error(res, "an error occurred", http.StatusBadRequest)
		}
		data.ToJson(res)
	}
	if req.Method == http.MethodGet {
		echo.getUsers(res, req)
		return
	}
	res.WriteHeader(http.StatusMethodNotAllowed)
}

func (echo *User) CreateUser(user *entity.User) (*entity.User, map[string]string) {
	return echo.app.SaveUser(user)
}

func (echo *User) GetUsers() (*entity.Users, error) {
	return echo.app.GetUsers()
}

func (echo *User) getUsers(res http.ResponseWriter, req *http.Request) {
	users, err := echo.app.GetUsers()
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}

	err = users.ToJson(res)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}

}

func createUser(res http.ResponseWriter, req *http.Request) *entity.User {
	userDto := dto.User{}
	err := userDto.FromJSON(req.Body)
	if err != nil {
		http.Error(res, "invalid entity", http.StatusBadRequest)
	}
	user := entity.User{
		Username: userDto.Username,
		Name:     userDto.Name,
	}
	return &user

}
