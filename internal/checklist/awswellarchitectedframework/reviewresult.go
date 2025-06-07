package awswellarchitectedframework

import "encoding/json"

type reviewResult struct {
	Title     string `json:"title"`
	Result    string `json:"result"`
	Reason    string `json:"reason"`
	Locations string `json:"locations"`
}

func ReviewResult(requirement Requirement, resp []byte) (string, error) {
	var result reviewResult
	if json.Unmarshal(resp, &result) != nil {
		result = reviewResult{
			Title:     string(requirement),
			Result:    "不明",
			Reason:    string(resp),
			Locations: "",
		}
	}
	b, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}
