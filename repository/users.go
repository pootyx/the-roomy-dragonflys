package repository

import (
	u "../utils"
	"fmt"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserId uint `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Password string `json:"password"`
	Email string `json:"email"`
	IsDeleted string `json:"is_deleted"`
	PictureUrl string `json:"picture_url"`
}

func GetUser(id uint) (*User) {
	contact := &User{}
	err := GetDB().Table("user").Where("id = ?", id).First(contact).Error
	if err != nil {
		return nil
	}
	return contact
}

func CreateUserToDb() {
	user := User{
		Model:      gorm.Model{},
		UserId:     1,
		FirstName:  "Teszt",
		LastName:   "Elek",
		Password:   "pasfksafas",
		Email:      "asasdas@gmail.com",
		IsDeleted:  "FALSES",
		PictureUrl: "https://c.disquscdn.com/uploads/users/16133/201/avatar92.jpg?1573664068",
	}

	GetDB().Create(&user)
	fmt.Println("User created!")
}

func (user *User) Create() (map[string]interface{}) {
	GetDB().Create(user)

	resp := u.Message(true, "success")
	resp["user"] = user
	return resp
}