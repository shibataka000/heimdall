package awswellarchitectedframework

import (
	"encoding/json"
	"fmt"
	"strings"
)

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
			Title:     strings.Split(string(requirement), " ")[0],
			Result:    "不明",
			Reason:    fmt.Sprintf("Amazon Bedrock の応答が不正です：%s", string(resp)),
			Locations: "",
		}
	}
	b, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}
