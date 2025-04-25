package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type UserRepository struct {
	connection *sql.DB
}

func NewUserRepository(connection *sql.DB) UserRepository{
	return UserRepository{
		connection: connection,
	}
}

func (user *UserRepository) GetUsers() ([]model.User, error){
	query := "SELECT id, nome, email, senha FROM usuarios"

	rows, err := user.connection.Query(query)
	if(err != nil){
		fmt.Println(err)
		return []model.User{}, err
	}

	var userList []model.User
	var userObj model.User

	for rows.Next(){
		err = rows.Scan(
			&userObj.ID,
			&userObj.Name,
			&userObj.Email,
			&userObj.Senha,
		)
		if(err != nil){
			fmt.Println(err)
			return []model.User{}, err
		}

		userList = append(userList, userObj)	
	}

	rows.Close()
	return userList, err
}

func (usr *UserRepository) CreateUser(user model.User) (int, error) {
	var id int
	query, err := usr.connection.Prepare("INSERT INTO usuarios " +
										 "(nome, email, senha) " +
										 "VALUES($1, $2, $3) RETURNING id")
	if(err != nil){
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(user.Name, user.Email, user.Senha).Scan(&id);
	if(err != nil){
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, nil
}