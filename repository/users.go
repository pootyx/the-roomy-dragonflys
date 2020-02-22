package repository

import (
	u "../utils"
	"fmt"
	"github.com/jinzhu/gorm"
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

func GetUsers() []*User {
	users := make([]*User, 0)
	err := GetDB().Find(users).Error
	if err != nil {
		return nil
	}
	return users
}

func GetUser(u string) *User {
	user := &User{}
	err := GetDB().Table("users").Where("user_id = ?", u).First(&user)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return user
}

func (user *User) Create() (map[string]interface{}) {
	GetDB().Create(user)

	resp := u.Message(true, "success")
	resp["user"] = user
	return resp
}