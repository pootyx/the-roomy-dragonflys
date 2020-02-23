package challenge_api

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/stack-attack/the-roomy-dragonflys/repository"
	"github.com/stack-attack/the-roomy-dragonflys/utils"
	"time"
)

type Challenge struct {
	gorm.Model
	UserId      string    `json:"userId"`
	ChallengeId string    `json:"challengeId"`
	Title       string    `json: "title"`
	Description string    `json:"description""`
	IsActive    bool      `json:"isActive"`
	EndDate     time.Time `json:"endDate"`
	Outcome     bool      `json:"outcome"`
	ProofUrl    string    `json:"proofUrl"`
}

func (challenge *Challenge) CreateChallenge() (map[string]interface{}, int) {
	challenge.ChallengeId = utils.GenerateUuid()
	challenge.IsActive = true

	if resp, valid := challenge.Validate(); !valid {
		return resp, 400
	}

	err := repository.GetDB().Create(challenge).Error
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	resp := utils.Message("success")
	resp["challenge"] = challenge
	return resp, 201
}

func (challenge *Challenge) UpdateChallenge(id string) (map[string]interface{}, int) {

	err := repository.GetDB().Table("challenges").Where("challenge_id = ?", id).Update(challenge)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	resp := utils.Message("success")
	resp["challenge"] = challenge
	return resp, 200

}

func GetChallengeById(id string) *Challenge {
	challenge := &Challenge{}
	err := repository.GetDB().Table("challenges").Where("challenge_id = ?", id).First(challenge).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return challenge
}

func GetChallengeByUserId(id string) []*Challenge {
	challenges := make([]*Challenge, 0)
	err := repository.GetDB().Table("challenges").Where("user_id = ?", id).Find(&challenges).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return challenges
}

func GetAllChallenge() []*Challenge {
	challenges := make([]*Challenge, 0)
	err := repository.GetDB().Table("challenges").Find(&challenges).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return challenges
}

func (challenge *Challenge) Validate() (map[string]interface{}, bool) {
	if challenge.UserId == "" {
		return utils.Message("UserId data attribute is missing!"), false
	}

	if challenge.Title == "" {
		return utils.Message("Title data attribute is missing!"), false
	}

	if challenge.Description == "" {
		return utils.Message("Description data attribute is missing!"), false
	}

	if challenge.EndDate.IsZero() {
		return utils.Message("EndDate data attribute is missing!"), false
	}

	return utils.Message("Everything was fine."), true
}
