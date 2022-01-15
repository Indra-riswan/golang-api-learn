package handler

import (
	"learn2/dto"
	"learn2/helper"
	"learn2/models"
	"learn2/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authhandler struct {
	authservice service.ServiceAuth
	jwtservice  service.JwtService
}

func NewAuthHandler(auth service.ServiceAuth, jwt service.JwtService) *authhandler {
	return &authhandler{auth, jwt}
}

func (h *authhandler) Login(ctx *gin.Context) {
	var LoginDto dto.DTOLogin
	errDto := ctx.ShouldBindJSON(&LoginDto)
	if errDto != nil {
		respons := helper.BuildErrorRespons("Failed Login", errDto.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, respons)
		return
	}
	authservice := h.authservice.VerifyCrendential(LoginDto.Email, LoginDto.Password)
	if v, ok := authservice.(models.User); ok {
		generatetoken := h.jwtservice.GenerateToken(strconv.FormatUint(uint64(v.ID), 10))
		v.Token = generatetoken
		respons := helper.BuildRespons(true, "Ok!", v)
		ctx.JSON(200, respons)
		return
	}
	respons := helper.BuildErrorRespons("Failed sent request", "invalid token", helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, respons)

}
func (h *authhandler) Register(ctx *gin.Context) {
	var registerDto dto.DTORigister
	errDto := ctx.ShouldBindJSON(&registerDto)
	if errDto != nil {
		respons := helper.BuildErrorRespons("Failed to Create an Account", errDto.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respons)
		return
	}
	if !h.authservice.DuplicateEmail(registerDto.Email) {
		respons := helper.BuildErrorRespons("Failed", "User Already Registered", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusConflict, respons)
		return
	} else {
		usercreate := h.authservice.CreateUser(registerDto)
		id := strconv.FormatUint(uint64(usercreate.ID), 10)
		token := h.jwtservice.GenerateToken(id)
		// token := h.jwtservice.GenerateToken(strconv.FormatUint(uint64(usercreate.ID), 10))
		usercreate.Token = token
		respons := helper.BuildRespons(true, "Ok!", usercreate)
		ctx.JSON(http.StatusCreated, respons)
	}

}
