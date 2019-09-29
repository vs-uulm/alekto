package trustScore

import (
	"github.com/ma-zero-trust-prototype/shared_lib/enum/fields"
	"github.com/ma-zero-trust-prototype/shared_lib/enum/vars"
	"github.com/ma-zero-trust-prototype/shared_lib/structs"
	"github.com/ma-zero-trust-prototype/trust_engine/data/data_variables"
	"github.com/ma-zero-trust-prototype/trust_engine/trustScore/subjectiveLogic"
	"github.com/ma-zero-trust-prototype/trust_engine/util"
)

/**
 * calculate the user and device opinions if there are
 * new information for this network agent since the last update
 */
func calcTrustScores(agent structs.NetworkAgent, scores *structs.TrustScores) {

	lastUpdated := scores.NetworkAgent.UpdatedAt
	userOpinion := getOpinionByKindAndId(fields.User, agent.User.Id, lastUpdated)
	deviceOpinion := getOpinionByKindAndId(fields.Device, agent.Device.Id, lastUpdated)

	scores.User.Score = subjectiveLogic.CalcProjectedProbability(userOpinion)
	scores.Device.Score = subjectiveLogic.CalcProjectedProbability(deviceOpinion)
	scores.NetworkAgent.Score = calcNetworkAgentTrustScore(userOpinion, deviceOpinion)

	return
}

func getOpinionByKindAndId(kind fields.Kind, id string, lastUpdated int64) structs.BinomialOpinion {

	subject := structs.Subject{
		Kind: kind,
		Name: id}

	variables := data_variables.GetVariablesBySubject(subject)

	if util.IsUnIdentifiedDevice(subject) { // get default values for unknown device
		return subjectiveLogic.CalcCumulativeFusedOpinion(getOpinionsOfVariables(variables))
	}

	for index, variable := range variables {

		variables[index] = updateOpinionForSubject(subject, variable, lastUpdated)
	}

	return subjectiveLogic.CalcCumulativeFusedOpinion(getOpinionsOfVariables(variables))
}

/**
 * calc network agent as cumulative fused opinion of user and device
 */
func calcNetworkAgentTrustScore(userOpinion, deviceOpinion structs.BinomialOpinion) structs.TrustScore {

	opinions := []structs.BinomialOpinion{userOpinion, deviceOpinion}
	agentOpinion := subjectiveLogic.CalcCumulativeFusedOpinion(opinions)

	return subjectiveLogic.CalcProjectedProbability(agentOpinion)
}

/**
 * update the subjects opinion according to recent information
 */
func updateOpinionForSubject(subject structs.Subject, variable structs.Variable,
	lastUpdated int64) (updatedVariable structs.Variable) {

	switch variable.Metadata.Name {

	case vars.UserAuthenticationAttempts:
		updatedVariable = updateUserByAuthenticationAttempts(subject, variable, lastUpdated)

	case vars.DeviceAuthenticationAttempts:
		updatedVariable = updateDeviceByAuthenticationAttempts(subject, variable, lastUpdated)
	}

	if variable.Opinion.HasDifferantValues(updatedVariable.Opinion) {
		data_variables.SaveNewVariablesBySubject(subject, updatedVariable)
	}

	return
}

func getOpinionsOfVariables(vars []structs.Variable) (opinions []structs.BinomialOpinion) {

	for _, variable := range vars {
		opinions = append(opinions, variable.Opinion)
	}

	return
}
