package storage

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type Storage struct{
	config *ConfigDb
	db *sql.DB
}


func NewStorage(db *ConfigDb) *Storage{
	return &Storage{config: db}
}


func (storage *Storage) OpenConnect() error{
	db, err := sql.Open("postgres", storage.config.DbUri)
	if err != nil{
		logrus.Printf("Error create db connection %s - %s", storage.config.DbUri, err)
		return err
	}
	err = db.Ping()
	if err != nil{
		logrus.Printf("Error open db connection %s - %s", storage.config.DbUri, err)
		return err
	}
	storage.db = db
	return nil
}

func (storage *Storage) CloseConnect() error{
	err := storage.db.Close()
	if err != nil{
		logrus.Println("Error close db connection ", err)
		return err
	}
	logrus.Println("Db connection close")
	return nil
}