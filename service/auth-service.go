package service

import (
	"fmt"
	"learn2/dto"
	"learn2/models"
	"learn2/repository"
	"log"

	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

type ServiceAuth interface {
	CreateUser(user dto.DTORigister) models.User
	VerifyCrendential(email string, password string) interface{}
	DuplicateEmail(email string) bool
	FindByEmail(email string) models.User
}

type serviceauth struct {
	repository repository.UserRepository
}

func NewServiceAuth(repository repository.UserRepository) *serviceauth {
	return &serviceauth{repository}
}

func (s *serviceauth) CreateUser(user dto.DTORigister) models.User {
	var users = models.User{}
	err := smapping.FillStruct(&users, smapping.MapFields(&user))
	if err != nil {
		log.Println(err)
		panic("Failed Create New User")
	}
	s.repository.InsertUser(users)
	return users

}
func ComparePassword(hashedPassword, plainpassword []byte) bool {
	p := []byte(hashedPassword)
	err := bcrypt.CompareHashAndPassword(p, plainpassword)
	if err != nil {
		fmt.Println("Incorrect Password", err.Error())
		return false
	}
	return true
}

func (s *serviceauth) VerifyCrendential(email string, password string) interface{} {
	res := s.repository.VerifyCredintial(email, password)
	if v, ok := res.(models.User); ok {
		compare := ComparePassword([]byte(v.Password), []byte(password))
		if v.Email == email && compare {
			return res
		}
		return false
	}
	return false
}

func (s *serviceauth) DuplicateEmail(email string) bool {
	res := s.repository.DuplicateEmail(email)
	return !(res.Error == nil)
}

func (s *serviceauth) FindByEmail(email string) models.User {
	return s.repository.FaindByEmail(email)

}
