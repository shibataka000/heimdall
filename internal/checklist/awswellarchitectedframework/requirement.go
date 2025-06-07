package awswellarchitectedframework

import (
	"encoding/json"
	"strings"

	_ "embed"

	_ "github.com/PuerkitoBio/goquery" // For genrequirements.go
)

//go:generate go run genrequirements.go
//go:embed requirements.json
var requirementsBytes []byte

// Requirement that design documents should satisfy.
// This is equivalent to the best practice in the AWS Well-Architected Framework.
type Requirement string

// Requirements returns the list of [Requirement] that design documents should satisfy.
// These are equivalent to the best practices in the AWS Well-Architected Framework.
func Requirements() ([]Requirement, error) {
	requirements := []Requirement{}
	if err := json.Unmarshal(requirementsBytes, &requirements); err != nil {
		return nil, err
	}
	return requirements, nil
}

// title returns the title of the [Requirement].
func (r Requirement) title() string {
	return strings.Split(string(r), " ")[0]
}
