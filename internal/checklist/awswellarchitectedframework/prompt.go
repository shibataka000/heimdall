package awswellarchitectedframework

import (
	"bytes"
	"text/template"

	_ "embed"
)

// Prompt for reviews based on the AWS Well-Architected Framework requirements.
type Prompt struct {
	requirement Requirement
}

//go:embed prompt.md
var promptTemplateBytes []byte

// NewPrompt creates a new [Prompt] for reviews based on the given AWS Well-Architected Framework requirements.
func NewPrompt(requirement Requirement) Prompt {
	return Prompt{requirement: requirement}
}

// Render generates the prompt text based on the AWS Well-Architected Framework requirement.
func (p Prompt) Render() (string, error) {
	promptTemplate, err := template.New("prompt").Parse(string(promptTemplateBytes))
	if err != nil {
		return "", err
	}
	var b bytes.Buffer
	data := map[string]string{
		"Requirement": string(p.requirement.Body),
	}
	if err := promptTemplate.Execute(&b, data); err != nil {
		return "", err
	}
	return b.String(), nil
}
