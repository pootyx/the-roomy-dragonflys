package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/stack-attack/the-roomy-dragonflys/bet-api"
	"github.com/stack-attack/the-roomy-dragonflys/challenge-api"
	"github.com/stack-attack/the-roomy-dragonflys/user-api"
	"os"
)

var db *gorm.DB //database

func init() {
	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=require password=%s", dbHost, username, dbName, dbPort, password)
	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.Debug().AutoMigrate(&user_api.User{}, &challenge_api.Challenge{}, &bet_api.Bet{})
}

func GetDB() *gorm.DB {
	return db
}