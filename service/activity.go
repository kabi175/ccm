package service

import (
	"errors"
	"time"

	"github.com/kabi175/ccm/http"
)

func Activity(handle string) (int, error) {
	submissions, err := http.GetSubmissions(handle)
	if err != nil {
		return 0, err
	}

	validated := RemoveDuplicate(submissions)
	DaySub := Filter(validated, time.Now().Add(time.Hour*-24))
	return len(DaySub), nil
}

func UpdateLevel(handle string) (*http.User, error) {
	user, err := http.GetLevel(handle)
	if user.Handle == "404" {
		return nil, errors.New("USer not found")
	}
	if err != nil {
		return nil, err
	}
	var min int
	switch user.Level {
	case 1:
		min = 3
	case 2:
		min = 4
	case 3:
		min = 5
	}
	progress, err := Activity(handle)
	if err != nil {
		return nil, err
	}
	if progress < min {
		user.Stars = 0
		return user, nil
	}
	user.Stars++
	if user.Stars >= 6 {
		user.Level += 1
		user.Stars = 0
	}
	return user, nil
}
