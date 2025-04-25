package usecase

import (
	"fmt"
	"go-api/model"
	"go-api/repository"
)

type UserUseCase struct {
	repository repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return UserUseCase{
		repository: repo,
	}
}

func (user *UserUseCase) GetUser() ([]model.User, error){
	return user.repository.GetUsers();
}

func (usr *UserUseCase) CreateUser(user model.User) (model.User, error){
	userID, err := usr.repository.CreateUser(user);
	if(err != nil){
		fmt.Println(err)
		return model.User{}, err
	}

	user.ID = userID
	return user, nil
}