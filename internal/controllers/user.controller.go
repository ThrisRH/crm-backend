package controllers

import (
	"go-sea-crm/internal/service"
	"go-sea-crm/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetUserInfo(c *gin.Context) {
	response.SuccessResponse(c, 20001, []string{"thrisx, tudun"})
	// response.ErrorResponse(c, 20003, "Not found!")
}
