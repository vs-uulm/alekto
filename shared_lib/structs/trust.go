package structs

import "github.com/ma-zero-trust-prototype/shared_lib/enum/fields"

type TrustScores struct {
	User         Score
	Device       Score
	NetworkAgent Score
}

type TrustScore = float32

type Score struct {
	Kind      fields.Kind `yaml:"kind"`
	Name      string      `yaml:"name"`
	UpdatedAt int64       `yaml:"updatedat"`
	Score     TrustScore  `yaml:"trustscore"`
}

func (score Score) Equals (other Score) bool {
	return score.Kind == other.Kind && score.Name == other.Name
}


type Variable struct {
	Kind        fields.Kind     `yaml:"kind"`
	Metadata    Metadata        `yaml:"metadata"`
	Subject     Subject         `yaml:"subject"`
	Opinion     BinomialOpinion `yaml:"opinion"`
}

func (variable Variable) IsSame (other Variable) bool {
	return variable.Subject.Equals(other.Subject) && variable.Metadata.Equals(other.Metadata)
}