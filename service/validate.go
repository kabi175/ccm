package service

import (
	"encoding/json"
	"io/ioutil"

	"github.com/kabi175/ccm/http"
)

type User struct {
	EmailId string `json:"email"`
	Handle  string `json:"Handle"`
}

const (
	filePermision = 0644
)

func Verify(filename string) ([]byte, []byte, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, nil, err
	}
	var (
		users        []User
		validUsers   []User
		inValidUsers []User
	)
	if err = json.Unmarshal(file, &users); err != nil {
		return nil, nil, err
	}
	for _, user := range users {
		err = http.VerifyProfile(user.Handle)
		if err != nil || user.Handle == "Errichto" || user.Handle == "Nil" || user.Handle == "nil" {
			inValidUsers = append(inValidUsers, user)
			continue
		}
		validUsers = append(validUsers, user)
	}

	var (
		valid   []byte
		invalid []byte
	)

	invalid, err = json.Marshal(inValidUsers)
	if err != nil {
		return nil, nil, err
	}

	valid, err = json.Marshal(validUsers)
	if err != nil {
		return nil, nil, err
	}

	return valid, invalid, err
}

func Extract(filename string) ([]byte, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var users []User
	if err = json.Unmarshal(file, &users); err != nil {
		return nil, err
	}

	var emails []string
	for _, user := range users {
		emails = append(emails, user.EmailId)
	}

	var content []byte
	content, err = json.Marshal(emails)
	if err != nil {
		return nil, err
	}
	return content, nil
}
