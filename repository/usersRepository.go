package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/stack-attack/the-roomy-dragonflys/utils"
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
		return nil
	}

	return user
}

func (user *User) CreateUser() (map[string]interface{}, int) {
	user.UserId = utils.GenerateUuid()

	if resp, valid := user.Validate(); !valid {
		return resp, 400
	}

	GetDB().Create(user)

	resp := utils.Message("success")
	resp["user"] = user
	return resp, 201
}

func DeleteUserById(userId string) {
	GetDB().Table("users").Where("user_id = ?", userId).Delete(&User{})
}

func (user *User) Validate() (map[string]interface{}, bool) {
	if user.UserId == "" {
		return utils.Message("UserId data attribute is missing!"), false
	}

	if user.FirstName == "" {
		return utils.Message("FirstName data attribute is missing!"), false
	}

	if user.LastName == "" {
		return utils.Message("LastName data attribute is missing!"), false
	}

	if user.Password == "" {
		return utils.Message("Password data attribute is missing!"), false
	}

	if user.Email == "" {
		return utils.Message("Email data attribute is missing!"), false
	}

	return utils.Message("Everything was fine."), true
}