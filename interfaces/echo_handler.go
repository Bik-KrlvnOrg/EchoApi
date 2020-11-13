package interfaces

import (
	"echoApi/app"
	"echoApi/domain/entity"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Echo struct {
	logger *log.Logger
	app    app.EchoAppInterface
}

func NewEcho(logger *log.Logger, appInterface app.EchoAppInterface) *Echo {
	return &Echo{logger, appInterface}
}

func (echo *Echo) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		item := createItem(res, req)
		data, err := echo.CreateItem(item)
		if err != nil {
			echo.logger.Println(err)
			http.Error(res, "an error occurred", http.StatusBadRequest)
		}
		data.ToJSON(res)
		return
	}

	if req.Method == http.MethodGet {
		echo.GetItems(res, req)
		return
	}

}

func (echo *Echo) CreateItem(item *entity.Item) (*entity.Item, error) {
	return echo.app.SaveEchos(item)
}

func (echo *Echo) GetItems(res http.ResponseWriter, req *http.Request) {
	items, err := echo.app.GetEchos()
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
	}
	err = items.ToJSON(res)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func createItem(res http.ResponseWriter, req *http.Request) *entity.Item {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(res, "an error occurred", http.StatusBadRequest)
	}
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		http.Error(res, "an error occurred", http.StatusBadRequest)
	}
	item := entity.Item{Echos: result}
	return &item
}
