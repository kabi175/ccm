package service

import (
	"fmt"
	"time"

	"github.com/kabi175/ccm/http"
)

func RemoveDuplicate(submissions []http.Submission) map[string]http.Submission {
	validated := make(map[string]http.Submission)
	for _, submission := range submissions {
		key := fmt.Sprint(submission.Problem.Id, submission.Problem.Index)
		if sub, ok := validated[key]; !ok || sub.Time > submission.Time {
			validated[key] = submission
		}
	}
	return validated
}

func Filter(submissions map[string]http.Submission, date time.Time) map[string]http.Submission {
	validated := make(map[string]http.Submission)
	for _, submission := range submissions {
		key := fmt.Sprint(submission.Problem.Id, submission.Problem.Index)
		subTime := time.Unix(submission.Time, 0)
		if subTime.Day() == date.Day() && subTime.Month() == date.Month() && subTime.Year() == date.Year() {
			validated[key] = submission
		}
	}
	return validated
}
