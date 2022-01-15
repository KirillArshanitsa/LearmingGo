package api

import (
	"StandartWebServer/storage"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (api *Api)ConfigLoggerFiled() error{
	logLevel, err := logrus.ParseLevel(api.config.LogLevel)
	if err != nil{
		logrus.Println("Error parse log level %s - %s", api.config.LogLevel, err)
		return err
	}
	logrus.SetLevel(logLevel)
	logrus.Println("Set log level", api.config.LogLevel)
	return nil
}


func (api *Api) configureRouterFiled(){
	api.router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("It's work!"))
	})
}

func (api *Api)configureDbFiled() *storage.Storage{
	return storage.NewStorage(api.config.ConfigDb)
}