package api

import (
	"StandartWebServer/internal/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Msg struct {
	StatusCode uint `json:"status_code"`
	Message string `json:"message"`
	IsError bool `json:"is_error"`
}

func (api *Api) CreateUser(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	user := models.Users{}
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil{
		api.logger.Info("Get bad json ", err)
		writer.WriteHeader(http.StatusBadRequest)
		msg := Msg{StatusCode: 400, Message: fmt.Sprintf("%s", err), IsError: true}
		json.NewEncoder(writer).Encode(msg)
		return
	}
	err = api.storage.GetUserRepository().CreateUser(&user)
	if err != nil{
		api.logger.Info("Error create user ", err)
		writer.WriteHeader(http.StatusInternalServerError)
		msg := Msg{StatusCode: 500, Message: fmt.Sprintf("%s", err), IsError: true}
		json.NewEncoder(writer).Encode(msg)
		return
	}
	api.logger.Info("User created")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(user)
}

func (api *Api) GetAllUsers(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	users, err := api.storage.GetUserRepository().GetAllUsers()
	if err != nil{
		api.logger.Info("Error get all users ", err)
		writer.WriteHeader(http.StatusInternalServerError)
		msg := Msg{StatusCode: 500, Message: fmt.Sprintf("%s", err), IsError: true}
		json.NewEncoder(writer).Encode(msg)
		return
	}
	api.logger.Info("User created")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(users)
}

func (api *Api) DeleteUser(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	idForDel, err  := strconv.ParseUint(mux.Vars(request)["id"], 10, 64)
	if err != nil{
		Msg := Msg{
			StatusCode: 400,
			Message: fmt.Sprintf("Bad request %s", err),
			IsError: true,
		}
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(Msg)
		return
	}
	isDeleted, err := api.storage.GetUserRepository().DeleteUser(idForDel)
	if err != nil{
		Msg := Msg{
			StatusCode: 501,
			Message: fmt.Sprintf("Bad request %s", err),
			IsError: true,
		}
		writer.WriteHeader(http.StatusNotImplemented)
		json.NewEncoder(writer).Encode(Msg)
		return
	}

	if !isDeleted{
		Msg := Msg{
			StatusCode: 501,
			Message: fmt.Sprintf("User with id - %d do not exist", idForDel),
			IsError: false,
		}
		writer.WriteHeader(http.StatusNotImplemented)
		json.NewEncoder(writer).Encode(Msg)
		return
	}

	api.logger.Info("Delete users ", idForDel)
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(idForDel)
}

func (api *Api) ChangeUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	user := models.Users{}
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil{
		api.logger.Info("Bad json ", err)
		Msg := Msg{
			StatusCode: 400,
			Message: fmt.Sprintf("Bad json %s", err),
			IsError: true,
		}
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(Msg)
		return
	}
	api.logger.Info("Try update user with id = ", user.Id)
	isUpd, err := api.storage.GetUserRepository().ChangeUser(&user)
	if err != nil{
		api.logger.Info("Error update user with id ",user.Id, err)
		Msg := Msg{
			StatusCode: 500,
			Message: fmt.Sprintf("Error update user %s", err),
			IsError: true,
		}
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(Msg)
		return
	}
	if !isUpd{
		api.logger.Info("User with id = ", user.Id, " not exist")
		Msg := Msg{
			StatusCode: 400,
			Message: fmt.Sprintf("User with id = %d not exist", user.Id),
			IsError: true,
		}
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(Msg)
		return
	}
	api.logger.Info("User with id = ", user.Id, " updated")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(user)
}