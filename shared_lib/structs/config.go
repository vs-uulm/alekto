package structs

import (
	"github.com/ma-zero-trust-prototype/shared_lib/enum/fields"
	"github.com/ma-zero-trust-prototype/shared_lib/enum/scope"
)

type Metadata struct {
	Name         string        `yaml:"name"`
	Scope        scope.Service `yaml:"scope"`
	Path         string        `yaml:"path,omitempty"`
	Exchangeable Exchangeable  `yaml:"exchangeable,omitempty"`
	Description  string        `yaml:"description"`
}
func (metadata Metadata) Equals (other Metadata) bool {
	return metadata.Name == other.Name
}

type Subject struct {
	Kind fields.Kind `yaml:"kind"`
	Name string      `yaml:"name"`
	Not  bool        `yaml:"not,omitempty"`
}

func (subject Subject) Equals (other Subject) bool {

	return subject.Kind == other.Kind && subject.Name == other.Name
}

type Exchangeable = []string
