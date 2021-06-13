package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/kabi175/ccm/http"
	"github.com/kabi175/ccm/service"
)

type handler struct{}

func (handler) Default() {
	fmt.Println("Invalid args")
}

func (handler) Extract(filename string) {

	filepath := fmt.Sprintf("./user-data/%s", filename)
	file, err := service.Extract(filepath)

	if err != nil {
		fmt.Println("Failed", err)
		return
	}

	filepath = "./user-data/extracted.json"

	err = ioutil.WriteFile(filepath, file, 0644)
	if err != nil {
		fmt.Println("Failed", err)
		return
	}

	fmt.Println("Extracted")
}

func (handler) Track() {

	filename := "./user-data/valid_users.json"
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
	fmt.Println("DB Updated")
}

func (handler) Init() {
	var check string
	fmt.Scanln(&check)
	if check != "yes" {
		return
	}

	filename := "./user-data/valid_users.json"
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
		return
	}

	var (
		users []service.User
		Users []http.User
	)

	err = json.Unmarshal(file, &users)

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

	fmt.Println("Users Added")
}

func (handler) Verify() {

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

	err = ioutil.WriteFile("./user-data/valid_users.json", valid, 0644)
	if err != nil {
		log.Println(err)
		return
	}

	err = ioutil.WriteFile("./user-data/invalid_users.json", invalid, 0644)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("Verification Completed")
}
