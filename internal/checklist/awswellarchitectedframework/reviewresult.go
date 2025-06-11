package awswellarchitectedframework

import (
	"encoding/json"
	"fmt"
)

// ReviewResult represents the result of a review against a requirement.
type ReviewResult struct {
	Title     string   `json:"title"`
	Result    string   `json:"result"`
	Reason    string   `json:"reason"`
	Locations []string `json:"locations"`
}

// NewReviewResult creates a new [ReviewResult] against a requirement.
// The `response` parameter is the response from Amazon Bedrock, which should be a JSON format.
func NewReviewResult(requirement Requirement, response []byte) (ReviewResult, error) {
	var result ReviewResult
	if err := json.Unmarshal(response, &result); err != nil {
		return ReviewResult{}, fmt.Errorf("the response from Amazon Bedrock Agent is invalid: %s", string(response))
	}
	result.Title = requirement.Title
	return result, nil
}

// String returns the string representation of the [ReviewResult].
func (r ReviewResult) String() string {
	b, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(b)
}
