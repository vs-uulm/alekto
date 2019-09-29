package data_scores

import (
	"github.com/ma-zero-trust-prototype/shared_lib/enum/fields"
	"github.com/ma-zero-trust-prototype/shared_lib/structs"
	"github.com/ma-zero-trust-prototype/shared_lib/time"
	"github.com/ma-zero-trust-prototype/trust_engine/util"
)

func GetScoreBySubject(subject structs.Subject) (storedScore structs.Score) {

	var scoreExists bool

	storedScore, scoreExists = getScoreBySubject(subject)

	if !scoreExists {
		storedScore = getDefaultScoreByKind(subject.Kind)
	}

	return
}

/**
 * get stored scores for user, device and agent
 * if no device id is given, the devices trust score is zero
 */
func GetScoresByNetworkAgent(agent structs.NetworkAgent) (scores structs.TrustScores) {

	user := structs.Subject{Kind: fields.User, Name: agent.User.Id}
	scores.User = GetScoreBySubject(user)

	device := structs.Subject{Kind: fields.Device, Name: agent.Device.Id}
	scores.Device = GetScoreBySubject(device)

	netAgent := structs.Subject{Kind: fields.NetworkAgent, Name: util.GetNetworkAgentsName(agent)}
	scores.NetworkAgent = GetScoreBySubject(netAgent)

	return
}

func SaveScoreByKindAndId(kind fields.Kind, id string, score structs.TrustScore) {

	newScore := structs.Score{
		Kind:      kind,
		Name:      id,
		Score:     score,
		UpdatedAt: time.NowTimestamp()}

	saveNewScore(newScore)
}
