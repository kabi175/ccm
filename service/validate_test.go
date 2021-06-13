package service

import (
	"io/ioutil"
	"testing"
)

func Test(t *testing.T) {
	t.Run("Profile Validation", func(t *testing.T) {
		valid, invalid, err := Verify("./user-data/data.json")
		if err != nil {
			t.Error(err)
			return
		}
		ioutil.WriteFile("valid_users.json", valid, 0644)
		ioutil.WriteFile("invalid_users.json", invalid, 0644)
	})
}
