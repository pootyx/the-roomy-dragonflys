package repository

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserId uint
	FirstName string
	LastName string
	Password string
	Email string
	IsDeleted string
	PictureUrl string
}

func GetUser(id uint) (*User) {

	contact := &User{}
	err := GetDb().Table("user").Where("id = ?", id).First(contact).Error
	if err != nil {
		return nil
	}
	return contact
}