package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/kabi175/ccm/http"
	"github.com/kabi175/ccm/service"
)

func handleTrack(filename string) {

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
		return
	}

	var (
		userHandles []service.User
		users       []http.User
	)

	err = json.Unmarshal(file, &userHandles)

	for _, userHandle := range userHandles {
		user, err := service.UpdateLevel(userHandle.Handle)
		if err != nil {
			log.Println(err)
			continue
		}

		usr := http.User{
			Handle: user.Handle,
			Level:  user.Level,
			Stars:  user.Stars,
		}
		users = append(users, usr)
	}
	http.UpdateDB(users)
}

func handleInit(filename string) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
		return
	}
	var users []service.User
	err = json.Unmarshal(file, &users)
	var Users []http.User
	for _, user := range users {
		usr := http.User{
			Handle: user.Handle,
			Level:  0,
			Stars:  0,
		}
		Users = append(Users, usr)
	}
	statusCode, err := http.UpdateDB(Users)
	if err != nil || statusCode != 200 {
		log.Println("error: ", err)
		return
	}
	fmt.Println("Updated")
}

func handleVerify() {
	filename := "./user-data/data.json"
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
		return
	}
	var userHandles []service.User
	err = json.Unmarshal(file, &userHandles)
	if err != nil {
		log.Println(err)
		return
	}
	valid, invalid, err := service.Verify(filename)
	if err != nil {
		log.Println(err)
		return
	}
	err = ioutil.WriteFile("valid_users.json", valid, 0644)
	if err != nil {
		log.Println(err)
		return
	}
	err = ioutil.WriteFile("invalid_users.json", invalid, 0644)
	if err != nil {
		log.Println(err)
		return
	}
}
