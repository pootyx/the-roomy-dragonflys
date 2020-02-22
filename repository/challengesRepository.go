package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/stack-attack/the-roomy-dragonflys/utils"
	"time"
)

type Challenge struct {
	gorm.Model
	UserId      uint      `json:"userId" gorm:"not null"`
	ChallengeId string    `json:"challengeId"`
	Title	string 		  `json: "title" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"`
	IsActive    bool      `json:"isActive"`
	EndDate     time.Time `json:"endDate"`
	Outcome     bool      `json:"outcome"`
	ProofUrl    string    `json:"proofUrl"`
}

func (challenge *Challenge) CreateChallenge() map[string]interface{} {
	challenge.ChallengeId = utils.GenerateUuid()

	if resp, ok := challenge.Validate(); !ok {
		fmt.Println("valami")
		return resp
	}

	GetDB().Create(challenge)

	resp := utils.Message("success")
	resp["challenge"] = challenge
	return resp
}

func GetChallengeById(id string) *Challenge {
	challenge := &Challenge{}
	err := GetDB().Table("challenges").Where("challenge_id = ?", id).First(challenge).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return challenge
}

func GetAllChallenge() []*Challenge {
	challenges := make([]*Challenge, 0)
	err := GetDB().Table("challenges").Find(&challenges).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return challenges
}

func (challenge *Challenge) Validate() (map[string]interface{}, bool) {
	fmt.Println("Validation")

	if challenge.Title == "" {
		return utils.Message("Contact name should be on the payload"), false
	}

	return utils.Message("Everything was fine."), true
}

