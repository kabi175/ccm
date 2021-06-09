package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestProfile(t *testing.T) {
	var profiles = []struct {
		handle  string
		isVaild bool
	}{
		{handle: "kabilan.ec19", isVaild: true},
		{handle: "dummy", isVaild: true},
		{handle: "umy", isVaild: false},
		{handle: "dflk", isVaild: false},
	}
	for _, profile := range profiles {
		testname := fmt.Sprintf("%s", profile.handle)
		t.Run(testname, func(t *testing.T) {
			err := VerifyProfile(profile.handle)
			if err != nil && profile.isVaild {
				t.Errorf("got %v, want %v", err, nil)
			}
			if err == nil && !profile.isVaild {
				t.Errorf("got %v, want %v", err, "Invalid Profle")
			}
		})
	}
}

func TestGetSubmissions(t *testing.T) {
	var handles = []string{"kabilan.ec19", "dummy"}
	for _, handle := range handles {

		testname := fmt.Sprintf("%s", handle)

		t.Run(testname, func(t *testing.T) {
			submissions, err := GetSubmissions(handle)
			if err != nil {
				t.Error(err)
			}
			fileName := fmt.Sprintf("./user-data/%s.json", handle)
			data, err := json.Marshal(submissions)
			if err != nil {
				t.Error(err)
			}
			err = ioutil.WriteFile(fileName, data, 0644)
			if err != nil {
				t.Error(err)
			}
		})
	}
}
