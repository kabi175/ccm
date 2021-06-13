package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

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
