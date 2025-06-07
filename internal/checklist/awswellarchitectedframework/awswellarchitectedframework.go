// Package awswellarchitectedframework provides functionality related to the AWS Well-Architected Framework requirements.
package awswellarchitectedframework

import (
	"bytes"
	"encoding/json"
	"text/template"

	_ "embed"

	_ "github.com/PuerkitoBio/goquery" // For generator.go.
)

//go:embed prompt.md
var promptTemplateBytes []byte

//go:generate go run generator.go
//go:embed requirements.json
var requirementsBytes []byte

// Requirement that design documents should satisfy.
// This is equivalent to the best practice in the AWS Well-Architected Framework.
type Requirement string

// Requirements returns requirements that design documents should satisfy.
// These are equivalent to the best practices in the AWS Well-Architected Framework.
func Requirements() ([]Requirement, error) {
	requirements := []Requirement{}
	if err := json.Unmarshal(requirementsBytes, &requirements); err != nil {
		return nil, err
	}
	return requirements, nil
}

// Prompt returns a prompt for reviews based on the given requirement in the AWS Well-Architected Framework.
func Prompt(requirement Requirement) (string, error) {
	promptTemplate, err := template.New("prompt").Parse(string(promptTemplateBytes))
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	data := map[string]string{
		"Requirement": string(requirement),
	}
	if err := promptTemplate.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

type ReviewResultX struct {
	Title     string `json:"title"`
	Result    string `json:"result"`
	Reason    string `json:"reason"`
	Locations string `json:"locations"`
}

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
