package data

import (
	"encoding/json"
	"fmt"
	"io"
)

type projects struct {
	Title       string
	Description string
	completed   bool
}

func GetProject(title string) *projects {

	_, pos, err := FindProject(title)
	if err != nil {
		fmt.Errorf("Invalid request Project not found")
	}
	return ProjectList[pos]

}

func (P *projects) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(P)
}

func (P *projects) ToDATA(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(P)
}

var err = fmt.Errorf("project not found")

func FindProject(title string) (*projects, int, error) {

	for i, p := range ProjectList {
		if p.Title == title {
			return p, i, nil
		}
	}
	return nil, -1, err
}

var ProjectList = []*projects{
	{
		Title:       "Java",
		Description: "Java API",
		completed:   false,
	},
	{
		Title:       "GO",
		Description: "GO API",
		completed:   true,
	},
}
