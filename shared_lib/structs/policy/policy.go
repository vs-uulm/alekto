package policy

import (
	"fmt"
	"github.com/ma-zero-trust-prototype/shared_lib/data/ip2location"
	"github.com/ma-zero-trust-prototype/shared_lib/structs"
)

type PolicyEvaluation struct {
	Success      bool
	Name         string
	Exchangeable structs.Exchangeable
	Message      string
}

type ClientPolicy struct {
	Kind     string            `yaml:"kind"`
	Metadata structs.Metadata  `yaml:"metadata"`
	Subjects []structs.Subject `yaml:"subjects"`
	Requires Requirements      `yaml:"requires"`
}

type Requirements struct {
	User         *User         `yaml:"user,omitempty"`
	Device       *Device       `yaml:"device,omitempty"`
	NetworkAgent *NetworkAgent `yaml:"networkagent,omitempty"`
}

type User struct {
	Authentication *StringArray `yaml:"authentication,omitempty"`
	TrustScore     *TrustScore  `yaml:"trustscore,omitempty"`
	Role           *StringArray `yaml:"role,omitempty"`
}

type Device struct {
	Authentication *StringArray          `yaml:"authentication,omitempty"`
	TrustScore     *TrustScore           `yaml:"trustscore,omitempty"`
	Type           *StringArray          `yaml:"type,omitempty"`
	Location       *ip2location.Location `yaml:"location,omitempty"`
}

type NetworkAgent struct {
	TrustScore *TrustScore `yaml:"trustscore,omitempty"`
}

// Requirements
type TrustScore = float32

type StringArray struct {
	Values   []string `yaml:"values"`
	Operator string   `yaml:"operator"`      // default or
	Not      bool     `yaml:"not,omitempty"` // default false
}

func (policy ClientPolicy) Print() {

	fmt.Printf("--- POLICY \n")
	fmt.Printf("%+v \n", policy.Metadata)
	for _, subject := range policy.Subjects {
		fmt.Printf("%+v \n", subject)
	}
}
