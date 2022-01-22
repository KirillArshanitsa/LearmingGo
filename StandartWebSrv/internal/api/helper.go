package api

import (
	"StandartWebServer/storage"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (api *Api)ConfigLoggerFiled() error{
	logLevel, err := logrus.ParseLevel(api.config.LogLevel)
	if err != nil{
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

	api.router.HandleFunc("/CreateUser", api.CreateUser).Methods("POST")
	api.router.HandleFunc("/GetAllUsers", api.GetAllUsers).Methods("GET")
	api.router.HandleFunc("/DeleteUser/{id}", api.DeleteUser).Methods("DELETE")
	api.router.HandleFunc("/ChangeUser", api.ChangeUser).Methods("POST")
}

func (api *Api)configureDbFiled() error{
	storageLocal := storage.NewStorage(api.config.ConfigDb)
	err := storageLocal.OpenConnect()
	if err != nil{
		return err
	}
	api.storage = storageLocal
	return nil
}