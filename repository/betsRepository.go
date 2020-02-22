package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Bet struct {
	gorm.Model
	// TODO Implement
}

func GetBetsByUserId(id string) *Bet {
	bet := &Bet{}
	err := GetDB().Table("bets").Where("user_id = ?", id).First(bet).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return bet
}
