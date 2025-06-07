package awswellarchitectedframework

import (
	"bytes"
	"text/template"

	_ "embed"
)

//go:embed prompt.md
var promptTemplateBytes []byte

// Prompt for reviews based on the AWS Well-Architected Framework requirements.
type Prompt string

// NewPrompt creates a new [Prompt] for reviews based on the given AWS Well-Architected Framework requirements.
func NewPrompt(requirement Requirement) (Prompt, error) {
	promptTemplate, err := template.New("prompt").Parse(string(promptTemplateBytes))
	if err != nil {
		return "", err
	}
	var b bytes.Buffer
	data := map[string]string{
		"Requirement": string(requirement),
	}
	if err := promptTemplate.Execute(&b, data); err != nil {
		return "", err
	}
	return Prompt(b.String()), nil
}

// String returns the string representation of the [Prompt].
func (p Prompt) String() string {
	return string(p)
}
