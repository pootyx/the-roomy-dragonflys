package repository

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/stack-attack/the-roomy-dragonflys/utils"
	u "github.com/stack-attack/the-roomy-dragonflys/utils"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strings"
)

/*
JWT claims struct
*/
type Token struct {
	UserId uint
	jwt.StandardClaims
}

type User struct {
	gorm.Model
	UserId     string `json:"userId"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	IsDeleted  string `json:"isDeleted"`
	PictureUrl string `json:"pictureUrl"`
	Token      string `json:"token";sql:"-"`
}

//Validate incoming user details...
func (user *User) Authorize() (map[string]interface{}, bool) {

	if !strings.Contains(user.Email, "@") {
		return u.Message("Email address is required"), false
	}

	if len(user.Password) < 6 {
		return u.Message("Password is required"), false
	}

	//Email must be unique
	temp := &User{}

	//check for errors and duplicate emails
	err := GetDB().Table("users").Where("email = ?", user.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message("Connection error. Please retry"), false
	}
	if temp.Email != "" {
		return u.Message("Email address already in use by another user."), false
	}

	return u.Message("Requirement passed"), true
}

func (user *User) Register() map[string]interface{} {

	user.UserId = utils.GenerateUuid()

	if resp, ok := user.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	GetDB().Create(user)

	if user.ID <= 0 {
		return u.Message("Failed to create user, connection error.")
	}

	//Create new JWT token for the newly registered account
	tk := &Token{UserId: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	user.Token = tokenString

	user.Password = "" //delete password

	response := u.Message("User has been created")
	response["user"] = user
	return response
}

func Login(email, password string) map[string]interface{} {

	user := &User{}
	err := GetDB().Table("users").Where("email = ?", email).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message("Email address not found")
		}
		return u.Message("Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return u.Message("Invalid login credentials. Please try again")
	}
	//Worked! Logged In
	user.Password = ""

	//Create JWT token
	tk := &Token{UserId: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	user.Token = tokenString //Store the token in the response

	resp := u.Message("Logged In")
	resp["user"] = user
	return resp
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
