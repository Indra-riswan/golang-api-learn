package repository

import (
	"learn2/models"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	InsertUser(user models.User) models.User
	UpdateUser(user models.User) models.User
	VerifyCredintial(email string, password string) interface{}
	FaindByEmail(email string) models.User
	DuplicateEmail(email string) (tx *gorm.DB)
	UserProfil(userID string) models.User
}

type userrepository struct {
	db *gorm.DB
}

func NewUserReposiroty(db *gorm.DB) *userrepository {
	return &userrepository{db}
}

func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed Generate Password")
	}
	return string(hash)

}

func (r *userrepository) InsertUser(user models.User) models.User {
	user.Password = HashAndSalt([]byte(user.Password))
	r.db.Save(&user)
	return user
}
func (r *userrepository) UpdateUser(user models.User) models.User {
	if user.Password != "" {
		user.Password = HashAndSalt([]byte(user.Password))
	} else {
		var tempuser models.User
		r.db.Find(tempuser, user.ID)
		user.Password = tempuser.Password
	}
	r.db.Save(&user)
	return user
}
func (r *userrepository) VerifyCredintial(email string, password string) interface{} {
	var user models.User
	res := r.db.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil

}
func (r *userrepository) FaindByEmail(email string) models.User {
	var user models.User
	r.db.Where("emial = ?", email).Take(&user)
	return user

}
func (r *userrepository) DuplicateEmail(email string) (tx *gorm.DB) {
	var user models.User
	return r.db.Where("email=?", email).Take(&user)

}
func (r *userrepository) UserProfil(userID string) models.User {
	var user models.User
	r.db.Find(&user, userID)
	return user

}
