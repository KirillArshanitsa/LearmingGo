package storage

import (
	"database/sql"
	"github.com/sirupsen/logrus"
)

type Storage struct{
	DbUri string
}


func NewStorage(db *ConfigDb) *Storage{
	return &Storage{DbUri: db.DbUri}
}


func (storage *Storage) OpenConnect() error{
	db, err := sql.Open("postgres", storage.DbUri)
	if err != nil{
		logrus.Println("Error create db connection %s - %s", storage.DbUri, err)
		return err
	}
	err = db.Ping()
	if err != nil{
		logrus.Println("Error open db connection %s - %s", storage.DbUri, err)
		return err
	}
	return nil
}

func (storage *Storage) CloseConnect(){
	storage.CloseConnect()
	logrus.Println("Db %s connection close", storage.DbUri)
}