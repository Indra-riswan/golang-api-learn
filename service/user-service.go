package service

import (
	"fmt"
	"learn2/dto"
	"learn2/models"
	"learn2/repository"

	"github.com/mashingan/smapping"
)

type UserService interface {
	Update(userdto dto.DTOUpdateUser) models.User
	UserProfil(ID string) models.User
}

type anakbabi struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *anakbabi {
	return &anakbabi{userRepo}
}

func (s *anakbabi) Update(userdto dto.DTOUpdateUser) models.User {
	var users = models.User{}
	err := smapping.FillStruct(&users, smapping.MapFields(&userdto))
	if err != nil {
		fmt.Println(err)

	}
	userUpdated := s.userRepo.UpdateUser(users)
	return userUpdated
}

func (s *anakbabi) UserProfil(ID string) models.User {
	return s.userRepo.UserProfil(ID)

}
