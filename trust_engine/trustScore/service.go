package trustScore

import (
	"fmt"
	"github.com/ma-zero-trust-prototype/shared_lib/enum/fields"
	"github.com/ma-zero-trust-prototype/shared_lib/structs"
	"github.com/ma-zero-trust-prototype/shared_lib/time"
	"github.com/ma-zero-trust-prototype/trust_engine/data/data_scores"
	"github.com/ma-zero-trust-prototype/trust_engine/util"
)

func CalcTrustScoresForAgent(agent structs.NetworkAgent) (agentScores structs.TrustScores) {

	if agent.Device.Id == "" {
		agent.Device.Id = util.GetUnIdentifiedDeviceName(agent)
	}

	agentScores = data_scores.GetScoresByNetworkAgent(agent)

	if updatedInTheLastTenSeconds(agentScores.NetworkAgent.UpdatedAt) {
		printScores(agentScores, false)
		return
	}

	calcTrustScores(agent, &agentScores)

	data_scores.SaveScoreByKindAndId(fields.User, agent.User.Id, agentScores.User.Score)
	data_scores.SaveScoreByKindAndId(fields.Device, agent.Device.Id, agentScores.Device.Score)
	data_scores.SaveScoreByKindAndId(fields.NetworkAgent, util.GetNetworkAgentsName(agent), agentScores.NetworkAgent.Score)

	printScores(agentScores, true)

	return
}


func updatedInTheLastTenSeconds (lastUpdated int64) bool {
	return (lastUpdated + 10) >= time.NowTimestamp()
}


func printScores (agentScores structs.TrustScores, new bool) {
	if new {
		fmt.Printf("New Calculated Scores ")
	} else {
		fmt.Printf("Loaded Scores ")
	}

	fmt.Printf("for Agent %s \n", agentScores.NetworkAgent.Name)
	fmt.Printf("User: %v | ", agentScores.User.Score)
	fmt.Printf("Device: %v | ", agentScores.Device.Score)
	fmt.Printf("Agent:  %v \n", agentScores.NetworkAgent.Score)
}