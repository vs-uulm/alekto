package subjectiveLogic

import (
	"github.com/ma-zero-trust-prototype/shared_lib/structs"
)

func CalcCumulativeFusedOpinion(opinions []structs.BinomialOpinion) (finalOpinion structs.BinomialOpinion) {

	for index, userOpinion := range opinions {

		if index == 0 {
			finalOpinion = userOpinion

		} else {
			finalOpinion = calcCumulativeFusedOpinion(userOpinion, finalOpinion)
		}
	}

	return
}

func CalcProjectedProbability(opinion structs.BinomialOpinion) float32 {

	return calcProjectedProbability(opinion)
}
