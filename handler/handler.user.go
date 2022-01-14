package handler

import (
	"fmt"
	"learn2/dto"
	"learn2/helper"
	"learn2/service"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	UpdateUser(ctx *gin.Context)
	UserProfil(ctx *gin.Context)
}

type userhandler struct {
	userService service.UserService
	jwtService  service.JwtService
}

func NewUserHandler(service service.UserService, jwt service.JwtService) UserHandler {
	return &userhandler{
		userService: service,
		jwtService:  jwt,
	}
}

func (h *userhandler) UpdateUser(ctx *gin.Context) {
	var userup dto.DTOUpdateUser
	err := ctx.ShouldBindJSON(&userup)
	if err != nil {
		respons := helper.BuildErrorRespons("Failed to Update", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, respons)
		return
	}

	authheader := ctx.GetHeader("Authorization")
	token, err := h.jwtService.ValidateToken(authheader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id, _ := strconv.Atoi(fmt.Sprintf("%v", claims["user_id"]))
	userup.ID = uint(id)
	up := h.userService.Update(userup)
	respons := helper.BuildRespons(true, "User Successed Updated", up)
	ctx.JSON(http.StatusAccepted, respons)
}

func (h *userhandler) UserProfil(ctx *gin.Context) {
	authheader := ctx.GetHeader("Authorization")
	token, err := h.jwtService.ValidateToken(authheader)
	if err != nil {
		panic(err.Error())
	}

	claims := token.Claims.(jwt.MapClaims)
	user := h.userService.UserProfil(fmt.Sprintf("%v", claims["user_id"]))
	res := helper.BuildRespons(true, "Ok!", user)
	ctx.JSON(200, res)
}
