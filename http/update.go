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
	Level  int    `json:"level"`
	Stars  int    `json:"stars"`
}

func UpdateDB(users []User) (int, error) {
	url := os.Getenv("POST_END_POINT")

	if url == "" {
		return 0, errors.New("Not a valid endpoint")
	}
	var userstr []UserStr
	for _, user := range users {
		userstr = append(userstr, conv(user))
	}
	postBody, err := json.Marshal(struct {
		Payload []UserStr `json:"payload"`
	}{Payload: userstr})

	if err != nil {
		return 0, err
	}

	responseBody := bytes.NewBuffer(postBody)

	res, err := http.Post(url, "application/json", responseBody)
	return res.StatusCode, err
}
