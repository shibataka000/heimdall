// Package awswellarchitectedframework provides functionality related to the AWS Well-Architected Framework requirements.
package awswellarchitectedframework

import (
	"bytes"
	"fmt"
	"net/http"
	"text/template"

	_ "embed"

	"github.com/PuerkitoBio/goquery"
)

//go:embed prompt.md
var promptTemplateBytes []byte

// GeneratePrompt creates a prompt for reviews based on AWS Well-Architected Framework requirements retrieved from the specified URL.
func GeneratePrompt(url string) (string, error) {
	promptTemplate, err := template.New("prompt").Parse(string(promptTemplateBytes))
	if err != nil {
		return "", err
	}
	requirement, err := fetchMainContent(url)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	data := map[string]string{
		"Requirement": requirement,
	}
	if err := promptTemplate.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func fetchMainContent(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close() // nolint:errcheck
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch content from %s: %s", url, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}
	return doc.Find("#main-content").Text(), nil
}
