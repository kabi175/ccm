package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Submission struct {
	Id      int    `json:"id"`
	Status  string `json:"verdict"`
	Problem struct {
		Id    int    `json:"contestId"`
		Index string `json:"index"`
	}
}

func GetSubmissions(userHandle string) ([]Submission, error) {
	type Responce struct {
		Result []Submission `json:"result"`
	}

	var (
		responce    Responce
		submissions []Submission
		index       int = 1
	)
	for {

		url := fmt.Sprintf(
			"https://codeforces.com/api/user.status?handle=%s&from=%d&count=1000", userHandle, index)

		rep, err := http.Get(url)

		if err != nil {
			return nil, err
		}

		defer rep.Body.Close()

		body, err := ioutil.ReadAll(rep.Body)

		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(body, &responce)

		if err != nil {
			return nil, err
		}

		submissions = append(submissions, responce.Result...)

		if len(responce.Result) < 1000 {
			return submissions, nil
		}
		index += 1000
	}
}

func VerifyProfile(userHandle string) error {
	url := fmt.Sprintf("https://codeforces.com/profile/%s", userHandle)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	if resp.Request.URL.String() == url {
		return nil
	}
	return errors.New("Invalid Profile")
}
