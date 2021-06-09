package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"
)

type User struct {
	Handle string `json:"handle"`
	Level  string `json:"level"`
}

func UpdateDB(users []User) (int, error) {
	url := os.Getenv("POST_END_POINT")

	if url == "" {
		return 0, errors.New("Not a valid end point")
	}

	postBody, err := json.Marshal(struct {
		Payload []User `json:"payload"`
	}{Payload: users})

	if err != nil {
		return 0, err
	}

	responseBody := bytes.NewBuffer(postBody)

	res, err := http.Post(url, "application/json", responseBody)
	return res.StatusCode, err
}
