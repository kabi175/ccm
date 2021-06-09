package http

import (
	"testing"

	"github.com/joho/godotenv"
)

func TestUpdate(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		t.Error(err)
	}
	testname := "users"
	users := []User{
		{
			Handle: "user5",
			Level:  "5",
		},
		{
			Handle: "user2",
			Level:  "2",
		},
		{
			Handle: "user3",
			Level:  "3",
		},
		{
			Handle: "user4",
			Level:  "4",
		},
	}
	t.Run(testname, func(t *testing.T) {
		status, err := UpdateDB(users)
		if err != nil || status != 200 {
			t.Error("unexpected result ", err, status)
		}
	})
}
