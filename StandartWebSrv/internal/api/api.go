package api

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Api struct{
	logger *logrus.Logger
	router *mux.Router
	config *Config
}

func NewApi(config *Config) *Api{
	return &Api{logger: logrus.New(), router: mux.NewRouter(), config: config}
}

func (api *Api) Start() error{
	err := api.ConfigLoggerFiled()
	if err != nil{
		return err
	}
	api.configureRouterFiled()
	logrus.Println("Server start on port " + api.config.Port)
	storage := api.configureDbFiled()
	err = storage.OpenConnect()
	if err != nil{
		return err
	}
	storage.CloseConnect()

	return http.ListenAndServe("localhost:" + api.config.Port, api.router)
}