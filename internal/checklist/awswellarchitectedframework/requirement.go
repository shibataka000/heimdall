package awswellarchitectedframework

import (
	"encoding/json"
	"fmt"

	_ "embed"

	_ "github.com/PuerkitoBio/goquery" // for genrequirements.go
)

type Requirement struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

//go:generate go run genrequirements.go
//go:embed requirements.json
var requirementsBytes []byte

func GetRequirement(id string) (Requirement, error) {
	requirements := []Requirement{}
	if err := json.Unmarshal(requirementsBytes, &requirements); err != nil {
		return Requirement{}, nil
	}
	for _, r := range requirements {
		if r.ID == id {
			return r, nil
		}
	}
	return Requirement{}, fmt.Errorf("requirement '%s' was not found", id)
}
