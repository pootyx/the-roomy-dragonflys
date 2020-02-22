package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/stack-attack/the-roomy-dragonflys/utils"
)

type Bet struct {
	gorm.Model
	BetId string `json:"betId"`
	UserId string `json:"userId"`
	ChallengeId string `json:"challengeId"`
	Bet bool `json:"bet"`
	Amount int `json:"amount"`
	Result int `json:"result"`
}

func GetBetById(betId string) *Bet {
	bet := &Bet{}
	err := GetDB().Table("bets").Where("bet_id =?", betId).First(bet).Error

	if err != nil {
		return nil
	}

	return bet
}

func GetAllBets() []*Bet {
	bet := make([]*Bet, 0)
	err := GetDB().Find(&bet).Error
	if err != nil {
		return nil
	}
	return bet
}

func (bet *Bet) CreateBet() (map[string]interface{}, int) {
	bet.BetId = utils.GenerateUuid()

	if resp, valid := bet.Validate(); !valid {
		return resp, 400
	}

	GetDB().Create(bet)

	resp := utils.Message("success")
	resp["bet"] = bet
	return resp, 201
}

func (bet *Bet) Validate() (map[string]interface{}, bool) {
	if bet.UserId == "" {
		return utils.Message("UserId data attribute is missing!"), false
	}

	if bet.ChallengeId == "" {
		return utils.Message("ChallengeId data attribute is missing!"), false
	}

	if bet.Amount == 0 {
		return utils.Message("Amount data attribute is missing!"), false
	}


	return utils.Message("Everything was fine."), true
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