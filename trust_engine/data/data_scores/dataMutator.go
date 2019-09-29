package data_scores

import (
	"fmt"
	"github.com/ma-zero-trust-prototype/shared_lib/structs"
	"github.com/ma-zero-trust-prototype/shared_lib/yaml"
)

func saveNewScore(newScore structs.Score) {

	var scores = getAllScores()
	var scoreExists bool
	var err error

	for index, score := range scores {

		if score.Equals(newScore) {

			scoreExists = true
			scores[index] = newScore
			break
		}
	}

	if !scoreExists {
		err = yaml.AppendStructToFile(getFilename(), []structs.Score{newScore})
	} else {
		err = yaml.SaveStructToFile(getFilename(), scores)
	}

	if err != nil {
		fmt.Printf("error while saving scores to file: %s", err.Error())
	}
}
