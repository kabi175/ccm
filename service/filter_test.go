package service

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/kabi175/ccm/http"
)

func TestFilter(t *testing.T) {
	handles := []string{"kabilan.ec19", "i_am_mpw"}
	for _, handle := range handles {
		testname := fmt.Sprintf(handle)
		t.Run(testname, func(t *testing.T) {
			submissions, err := http.GetSubmissions(handle)
			log.Println("submissions ", len(submissions))
			if err != nil {
				t.Error(err)
				return
			}
			validated := RemoveDuplicate(submissions)
			log.Println("validated ", len(validated))
			filtered := Filter(validated, time.Now())
			log.Println("filtered ", len(filtered))
		})
	}
}
