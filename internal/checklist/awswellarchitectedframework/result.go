package awswellarchitectedframework

import "encoding/json"

type reviewResult struct {
	Title     string `json:"title"`
	Result    string `json:"result"`
	Reason    string `json:"reason"`
	Locations string `json:"locations"`
}

func ReviewResult(resp []byte) (string, error) {
	var result reviewResult
	if json.Unmarshal(resp, &result) != nil {
		result = reviewResult{
			Title:     "Review Result",
			Result:    "Unknown",
			Reason:    "The response is not in the expected format.",
			Locations: "N/A",
		}
	}
	b, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}
