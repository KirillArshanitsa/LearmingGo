package api

import (
	"StandartWebServer/storage"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Api struct{
	logger *logrus.Logger
	router *mux.Router
	config *Config
	storage *storage.Storage
}

func NewApi(config *Config) *Api{
	return &Api{logger: logrus.New(), router: mux.NewRouter(), config: config}
}

func (api *Api) Start() error{
	api.logger.Info("Configure logger")
	err := api.ConfigLoggerFiled()
	if err != nil{
		api.logger.Info("Error parse log level ", api.config.LogLevel, err)
		return err
	}
	api.logger.Info("Configure router")
	api.configureRouterFiled()
	api.logger.Info("Server start on port " + api.config.Port)
	api.logger.Info("Configure storage")
	err = api.configureDbFiled()
	if err != nil {
		api.logger.Info("Error configure storage", err)
		return err
	}
	api.logger.Info("Try to start web server")
	return http.ListenAndServe("localhost:" + api.config.Port, api.router)
}