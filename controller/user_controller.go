package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct {
	UserUseCase usecase.UserUseCase
}

func NewUserController(usecase usecase.UserUseCase) userController {
	return userController{
		UserUseCase: usecase,
	}
}

func (user *userController) GetUsers(ctx *gin.Context){
	users, err := user.UserUseCase.GetUser()
	if(err != nil){
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, users)
}

func (usr *userController) CreateUser(ctx *gin.Context){
	var user model.User
	err := ctx.BindJSON(&user)
	if(err != nil){
		ctx.JSON(http.StatusBadRequest, err)
		return	
	}

	insertedProduct, err := usr.UserUseCase.CreateUser(user)

	if(err != nil){
		ctx.JSON(http.StatusInternalServerError, err)
		return 
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}