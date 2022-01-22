package storage

import (
	"StandartWebServer/internal/models"
	"fmt"
)

type UserRepository struct {
	storage *Storage
}


var tableUser = "Users"


func (u *UserRepository) CreateUser(user *models.Users) error{
	sql := fmt.Sprintf("insert into %s (login,password) values($1,$2) returning id", tableUser)
	err := u.storage.db.QueryRow(sql, user.Login, user.Password).Scan(&user.Id)
	if err != nil{
		return err
	}
	return nil
}

func (u *UserRepository) GetAllUsers() ([]*models.Users, error){
	sql := fmt.Sprintf("select * from users")
	rows, err := u.storage.db.Query(sql)
	if err != nil{
		return nil, err
	}
	users := make([]*models.Users, 0)
	defer rows.Close()
	for rows.Next(){
		user := models.Users{}
		err = rows.Scan(&user.Id, &user.Login, &user.Password)
		if err != nil{
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (u *UserRepository) DeleteUser(id uint64)  (bool, error){
	sql := fmt.Sprintf("delete from %s where id=$1 returning id", tableUser)
	resp, err := u.storage.db.Exec(sql, id)
	if err != nil{
		return false, err
	}
	count, err :=resp.RowsAffected()
	if err != nil{
		return false ,err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}

func(u *UserRepository) ChangeUser(user *models.Users) (bool, error){
	sql := fmt.Sprintf("update %s set login=$1, password=$2 where id = $3 returning id", tableUser)
	resp, err := u.storage.db.Exec(sql, user.Login, user.Password, user.Id)
	if err != nil{
		return false ,err
	}
	count, err :=resp.RowsAffected()
	if err != nil{
		return false ,err
	}
	if count == 1 {
		return true, nil
	}
	return false, nil
}