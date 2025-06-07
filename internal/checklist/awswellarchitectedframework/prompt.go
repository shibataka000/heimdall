// Package awswellarchitectedframework provides functionality related to the AWS Well-Architected Framework requirements.
package awswellarchitectedframework

import (
	"bytes"
	"text/template"

	_ "embed"

	_ "github.com/PuerkitoBio/goquery" // For generator.go.
)

//go:embed prompt.md
var promptTemplateBytes []byte

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
