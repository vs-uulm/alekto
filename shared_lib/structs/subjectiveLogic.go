package structs

/**
 * Subjective Logic - A Formalism for Reasoning Under Uncertainty
 * Page 24
 */
type BinomialOpinion struct {
	B float32 `yaml:"belief"`
	D float32 `yaml:"disbelief"`
	U float32 `yaml:"uncertainty"`
	A float32 `yaml:"baserate"`
}

func (opinion BinomialOpinion) HasDifferantValues (other BinomialOpinion) bool {
	return opinion.D != other.D || opinion.B != other.B || opinion.U != other.U
}


func (opinion BinomialOpinion) RaiseBelief (percent float32) BinomialOpinion {

	if opinion.B == 1 {
		return opinion
	}

	newBelief := getRaisedValue(opinion.B, percent)
	diff := newBelief - opinion.B
	newDisbelief := betweenZeroAndOne(opinion.D - (diff / 2))
	newUncertainty := betweenZeroAndOne(1 - newBelief - newDisbelief)

	return BinomialOpinion{
		B: newBelief,
		D: newDisbelief,
		U: newUncertainty,
		A: opinion.A}
}


func (opinion BinomialOpinion) RaiseDisbelief (percent float32) BinomialOpinion {

	if opinion.D == 1 {
		return opinion
	}

	newDisbelief := getRaisedValue(opinion.D, percent)
	diff := newDisbelief - opinion.D
	newBelief := betweenZeroAndOne(opinion.B - (diff / 2))
	newUncertainty := betweenZeroAndOne(1 - newBelief - newDisbelief)

	return BinomialOpinion{
		B: newBelief,
		D: newDisbelief,
		U: newUncertainty,
		A: opinion.A}
}


func getRaisedValue (value float32, percent float32) float32 {
	if value == 0 {
		return percent
	} else {
		return betweenZeroAndOne(value * (percent + 1))
	}
}


func betweenZeroAndOne (value float32) float32 {

	if value > 1 {
		return 1
	}

	if value < 0 {
		return 0
	}

	return value
}