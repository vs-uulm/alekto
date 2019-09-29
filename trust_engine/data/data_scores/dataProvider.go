package data_scores

import (
	"flag"
	"github.com/ma-zero-trust-prototype/shared_lib/enum/fields"
	"github.com/ma-zero-trust-prototype/shared_lib/structs"
	"github.com/ma-zero-trust-prototype/shared_lib/yaml"
)

func getAllScores() (scores []structs.Score) {

	yaml.LoadStructFromFile(getFilename(), &scores)
	return
}

func getScoreBySubject(subject structs.Subject) (storedScore structs.Score, scoreExists bool) {

	var searchedScore = structs.Score{Kind: subject.Kind, Name: subject.Name}

	for _, score := range getAllScores() {

		if score.Equals(searchedScore) {

			storedScore, scoreExists = score, true
			break
		}
	}

	return
}

func getDefaultScoreByKind(kind fields.Kind) (score structs.Score) {

	var scores []structs.Score

	yaml.LoadStructFromFile(getFilenameDefault(), &scores)

	for _, score := range scores {

		if score.Kind == kind {

			return score
		}
	}

	return score
}

func getFilename() string {
	filename := "./data/data_scores/scores.yaml"

	if flag.Lookup("test.v") == nil {
		return filename
	} else {
		return "." + filename
	}
}

func getFilenameDefault() string {
	filename := "./data/data_scores/default.yaml"

	if flag.Lookup("test.v") == nil {
		return filename
	} else {
		return "." + filename
	}
}
