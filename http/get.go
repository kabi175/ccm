package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type UserStr struct {
	Handle string `json:"handle"`
	Level  string `json:"level"`
	Stars  string `json:"stars"`
}

func GetLevel(handle string) (*User, error) {
	url := os.Getenv("GET_END_POINT")
	if url == "" {
		return nil, errors.New("Not a valid endpoint")
	}
	url = fmt.Sprintf("%s%s", url, handle)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var user UserStr
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &user)
	level, err := strconv.Atoi(user.Level)
	stars, err := strconv.Atoi(user.Stars)
	return &User{
		Handle: user.Handle,
		Level:  level,
		Stars:  stars,
	}, nil
}

func conv(user User) UserStr {
	level := strconv.Itoa(user.Level)
	stars := strconv.Itoa(user.Stars)
	return UserStr{
		Handle: user.Handle,
		Level:  level,
		Stars:  stars,
	}
}
