package handler

import (
	"fmt"
	"learn2/dto"
	"learn2/helper"
	"learn2/models"
	"learn2/service"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type BookHandler interface {
	BookAll(ctx *gin.Context)
	FindBookById(ctx *gin.Context)
	CreateBook(ctx *gin.Context)
	UpdateBook(ctx *gin.Context)
	DeleteBook(ctx *gin.Context)
}

type bookhandler struct {
	serviceBook service.BookService
	serviceJwt  service.JwtService
}

func NewBookHandler(sBook service.BookService, sJwt service.JwtService) *bookhandler {
	return &bookhandler{sBook, sJwt}
}

func (h *bookhandler) getUserIDByToken(token string) string {
	bToken, err := h.serviceJwt.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}

	claims := bToken.Claims.(jwt.MapClaims)
	return fmt.Sprintf("%v", claims["user_id"])
}

func (h *bookhandler) BookAll(ctx *gin.Context) {
	var books []models.Book = h.serviceBook.AllBook()
	res := helper.BuildRespons(true, "OK!", books)
	ctx.JSON(200, res)
}

func (h *bookhandler) FindBookById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	book := h.serviceBook.FindBookById(uint(id))
	if (book == models.Book{}) {
		res := helper.BuildErrorRespons("No Data Found", "Invalid Id", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildRespons(true, "Ok!", book)
		ctx.JSON(200, res)
	}

}
func (h *bookhandler) CreateBook(ctx *gin.Context) {
	var bookDtoC dto.DTOBookregister
	errDTB := ctx.ShouldBindJSON(&bookDtoC)
	if errDTB != nil {
		res := helper.BuildErrorRespons("failed create book", errDTB.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	} else {
		autheader := ctx.GetHeader("Authorization")
		userID := h.getUserIDByToken(autheader)
		userIDc, _ := strconv.Atoi(userID)
		bookDtoC.UserID = uint(userIDc)
	}
	result := h.serviceBook.CreateBook(bookDtoC)
	res := helper.BuildRespons(true, "Ok!", result)
	ctx.JSON(http.StatusCreated, res)

}

func (h *bookhandler) UpdateBook(ctx *gin.Context) {
	var bookDtoU dto.DTOBookUpdate
	errDTB := ctx.ShouldBindJSON(&bookDtoU)
	if errDTB != nil {
		res := helper.BuildErrorRespons("failed update book", errDTB.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	authheader := ctx.GetHeader("Authorization")
	token, err := h.serviceJwt.ValidateToken(authheader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	if h.serviceBook.AlllowEdited(userID, bookDtoU.ID) {
		userIDc, _ := strconv.Atoi(userID)
		bookDtoU.ID = uint(userIDc)
		res := h.serviceBook.UpdateBook(bookDtoU)
		respons := helper.BuildRespons(true, "ok!", res)
		ctx.JSON(200, respons)
	} else {
		respons := helper.BuildErrorRespons("Permisson Danied", "You not Owner of the Book", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusForbidden, respons)
	}

}

func (h *bookhandler) DeleteBook(ctx *gin.Context) {
	var book models.Book
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		respons := helper.BuildErrorRespons("Failed Delete Book ", "No Id Found", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respons)
	}
	book.ID = uint(id)
	authheader := ctx.GetHeader("Authorization")
	token, err := h.serviceJwt.ValidateToken(authheader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if h.serviceBook.AlllowEdited(userID, book.ID) {
		h.serviceBook.DeleteBook(book)
		res := helper.BuildRespons(true, "Delete", helper.EmptyObj{})
		ctx.JSON(200, res)
	} else {
		res := helper.BuildErrorRespons("You dont have permission", "permission danied", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusForbidden, res)
	}

}
