package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
) 

func main(){
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if(err != nil){
		panic((err))
	}

	UserRepository := repository.NewUserRepository(dbConnection)
	UserUseCase := usecase.NewUserUseCase(UserRepository)
	UserController := controller.NewUserController(UserUseCase)

	server.GET("/users", UserController.GetUsers )
	server.POST("/user", UserController.CreateUser)
	server.Run(":8000")
}