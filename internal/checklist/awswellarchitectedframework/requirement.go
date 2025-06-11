package awswellarchitectedframework

import (
	"encoding/json"

	_ "embed"

	_ "github.com/PuerkitoBio/goquery" // For genrequirements.go
)

// Requirement that design documents should satisfy.
// This is equivalent to the best practice in the AWS Well-Architected Framework.
type Requirement struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

// Requirements that design documents should satisfy.
// These are equivalent to the best practices in the AWS Well-Architected Framework.
var Requirements []Requirement

//go:generate go run genrequirements.go
//go:embed requirements.json
var requirementsBytes []byte

func init() {
	if err := json.Unmarshal(requirementsBytes, &Requirements); err != nil {
		panic(err)
	}
}
