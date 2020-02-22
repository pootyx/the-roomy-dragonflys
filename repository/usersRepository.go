package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	util "github.com/stack-attack/the-roomy-dragonflys/utils"
)

type User struct {
	gorm.Model
	UserId string `json:"userId"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Password string `json:"password"`
	Email string `json:"email"`
	IsDeleted string `json:"isDeleted"`
	PictureUrl string `json:"pictureUrl"`
}

func GetAllUsers() []*User {
	users := make([]*User, 0)
	err := GetDB().Find(&users).Error
	if err != nil {
		return nil
	}
	return users
}

func GetUserById(userId string) *User {
	user := &User{}
	err := GetDB().Table("users").Where("user_id = ?", userId).First(user).Error

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return user
}

func (user *User) CreateUser() (map[string]interface{}) {
	user.UserId = util.GenerateUuid()
	fmt.Println(user.UserId)
	GetDB().Create(user)

	resp := util.Message("success")
	resp["user"] = user
	return resp
}