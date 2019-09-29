/**
 * Source:
 * Subjective Logic - A Formalism for Reasoning Under Uncertainty
 * Audun Jøsang
 */

package subjectiveLogic

import (
	"github.com/ma-zero-trust-prototype/shared_lib/structs"
)

/**
* Page 24
 */
func calcProjectedProbability(opinion structs.BinomialOpinion) float32 {

	return opinion.B + opinion.A*opinion.U
}

/**
 * In case of dogmatic arguments, the average of the values is calculated
 * Page 226-227
 */
func calcCumulativeFusedOpinion(op1, op2 structs.BinomialOpinion) (opinion structs.BinomialOpinion) {

	if op1.U == 0 && op2.U == 0 {

		opinion = cumulativeFusionOperatorDogmatic(op1, op2)

	} else {

		opinion = cumulativeFusionOperator(op1, op2)
	}

	return
}

/**
 * Subjective Logic - A Formalism for Reasoning Under Uncertainty
 * Page 226-227
 */
func cumulativeFusionOperatorDogmatic(op1, op2 structs.BinomialOpinion) (opinion structs.BinomialOpinion) {

	opinion.B = (op1.B + op2.B) / 2

	opinion.U = 0

	opinion.A = (op1.A + op2.A) / 2

	return
}

/**
 * Subjective Logic - A Formalism for Reasoning Under Uncertainty
 * Page 226
 */
func cumulativeFusionOperator(op1, op2 structs.BinomialOpinion) (opinion structs.BinomialOpinion) {

	opinion.B = (op1.B*op2.U + op2.B*op1.U) / (op1.U + op2.U - op1.U*op2.U)

	opinion.U = (op1.U * op2.U) / (op1.U + op2.U - op1.U*op2.U)

	if op1.U != 1 || op2.U != 1 {

		opinion.A = (op1.A*op2.U + op2.A*op1.U - (op1.A+op2.A)*op1.U*op2.U) /
			(op1.U + op2.U - 2*op1.U*op2.U)

	} else {

		opinion.A = (op1.A + op2.A) / 2
	}

	return
}
