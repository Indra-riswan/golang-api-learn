package main

import (
	"learn2/config"
	"learn2/handler"
	"learn2/models"
	"learn2/repository"
	"learn2/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.ConnectionDb()
	userRepository repository.UserRepository = repository.NewUserReposiroty(db)
	bookRepository repository.BookRepository = repository.NewBookRepository(db)
	authService    service.ServiceAuth       = service.NewServiceAuth(userRepository)
	bookService    service.BookService       = service.NewBookService(bookRepository)
	jwtJwt         service.JwtService        = service.NewJwtService()
	authHandler    handler.AuthHandler       = handler.NewAuthHandler(authService, jwtJwt)
	bookHandler    handler.BookHandler       = handler.NewBookHandler(bookService, jwtJwt)
	userHandler    handler.UserHandler       = handler.NewUserHandler(userService, jwtJwt)
	userService    service.UserService       = service.NewUserService(userRepository)
)

func main() {
	defer config.CloseConnectionDb(db)
	db.AutoMigrate(&models.Book{}, &models.User{})

	r := gin.Default()

	auth := r.Group("v2/auth")
	{
		auth.POST("/login", authHandler.Login)
		auth.POST("/register", authHandler.Register)

	}

	book := r.Group("v2/books", handler.AuthorizeJwt(jwtJwt))
	{
		book.POST("/create", bookHandler.CreateBook)
		book.GET("/all", bookHandler.BookAll)
		book.GET("/:id", bookHandler.FindBookById)
		book.PUT("/:id", bookHandler.UpdateBook)
		book.DELETE("/:id", bookHandler.DeleteBook)
	}

	user := r.Group("v2/users", handler.AuthorizeJwt(jwtJwt))

	user.PUT("/update", userHandler.UpdateUser)
	user.GET("/profil/:id", userHandler.UserProfil)

	r.Run()

}
