package repository

import (
	u "../utils"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"time"
)

type Challenge struct {
	gorm.Model
	UserId      uint      `json:"userId"`
	ChallengeId string    `json:"challengeId"`
	Description string    `json:"description"`
	IsActive    bool      `json:"isActive"`
	EndDate     time.Time `json:"endDate"`
	Outcome     bool      `json:"outcome"`
	ProofUrl    string    `json:"proofUrl"`
}

func (challenge *Challenge) Create() map[string]interface{} {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil
	}
	challenge.ChallengeId = uuid.String()

	GetDB().Create(challenge)

	resp := u.Message(true, "success")
	resp["challenge"] = challenge
	return resp
}

func GetChallenge(id string) *Challenge {
	challenge := &Challenge{}
	GetDB().Table("challenges").Where("challenge_id = ?", id).First(challenge)
	if challenge.Description == "" {
		return nil
	}
	return challenge
}
