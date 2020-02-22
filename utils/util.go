package utils

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
)

func Message(message string) (map[string]interface{}) {
	return map[string]interface{}{"message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func GenerateUuid() string {
	uuid, err := uuid.NewRandom()

	if err != nil {
		panic("UUID ERROR!")
	}

	return uuid.String()
}
