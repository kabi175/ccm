package http

import (
	"errors"
	"fmt"
	"net/http"
)

func VerifyProfile(userHandle string) error {
	url := fmt.Sprintf("https://codeforces.com/profile/%s", userHandle)
	var (
		resp *http.Response
		err  error
	)
	resp, err = http.Get(url)
	if err != nil {
		return err
	}
	if resp.Request.URL.String() != "https://codeforces.com/" && resp.StatusCode == 200 {
		return nil
	}
	return errors.New("Invalid Profile")
}
