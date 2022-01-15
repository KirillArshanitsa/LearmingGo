package api

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Api struct{
	Logger *logrus.Logger
	Router *mux.Router
	Config *Config
}

func NewApi(config *Config) *Api{
	return &Api{Logger: logrus.New(), Router: mux.NewRouter(), Config: config}
}

func (api *Api) Start() error{
	err := api.ConfigLoggerFiled()
	if err != nil{
		return err
	}
	api.configureRouterFiled()

	storage := api.configureDbFiled()
	err = storage.OpenConnect()
	if err != nil{
		return err
	}
	storage.CloseConnect()

	return nil
}