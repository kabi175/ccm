package http

import (
	"fmt"
	"testing"
)

func TestProfile(t *testing.T) {
	var profiles = []struct {
		handle  string
		isVaild bool
	}{
		{handle: "kabilan.ec19", isVaild: true},
		{handle: "dummy", isVaild: true},
		{handle: "umy", isVaild: true},
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
