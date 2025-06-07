package awswellarchitectedframework

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ReviewResult represents the result of a review against a requirement.
type ReviewResult struct {
	Title     string `json:"title"`
	Result    string `json:"result"`
	Reason    string `json:"reason"`
	Locations string `json:"locations"`
}

// NewReviewResult creates a new [ReviewResult] against a requirement.
// The `response` parameter is the response from Amazon Bedrock, which should be a JSON format.
func NewReviewResult(requirement Requirement, response []byte) ReviewResult {
	var result ReviewResult
	if json.Unmarshal(response, &result) != nil {
		result = newReviewResultForInvalidResponse(requirement, response)
	}
	return result
}

// String returns the string representation of the [ReviewResult].
func (r ReviewResult) String() string {
	b, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(b)
}

// newReviewResultForInvalidResponse creates a new [ReviewResult] for an invalid response from Amazon Bedrock.
func newReviewResultForInvalidResponse(requirement Requirement, response []byte) ReviewResult {
	return ReviewResult{
		Title:     strings.Split(string(requirement), " ")[0],
		Result:    "不明",
		Reason:    fmt.Sprintf("Amazon Bedrock の応答が不正です：%s", string(response)),
		Locations: "",
	}
}
