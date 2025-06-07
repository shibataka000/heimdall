package awswellarchitectedframework

import (
	"bytes"
	"text/template"

	_ "embed"
)

//go:embed prompt.md
var promptTemplateBytes []byte

// Prompt returns a prompt for reviews based on the given requirement in the AWS Well-Architected Framework.
func Prompt(requirement Requirement) (string, error) {
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
	return b.String(), nil
}
