package api

import (
	"StandartWebServer/storage"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (api *Api)ConfigLoggerFiled() error{
	logLevel, err := logrus.ParseLevel(api.Config.LogLevel)
	if err != nil{
		logrus.Println("Error parse log level %s - %s", api.Config.LogLevel, err)
		return err
	}
	logrus.SetLevel(logLevel)
	logrus.Println("Set log level -%s", api.Config.LogLevel)
	return nil
}


func (api *Api) configureRouterFiled(){
	api.Router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("It's work!"))
	})
}

func (api *Api)configureDbFiled() *storage.Storage{
	return storage.NewStorage(api.Config.ConfigDb)
}