package awswellarchitectedframework

import "encoding/json"

func ReviewResult(resp []byte) (string, error) {
	var result ReviewResultX
	if json.Unmarshal(resp, &result) != nil {
		result = ReviewResultX{
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
