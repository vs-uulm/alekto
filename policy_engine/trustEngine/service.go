package trustEngine

import (
	"fmt"
	"github.com/ma-zero-trust-prototype/policy_engine/env"
	"github.com/ma-zero-trust-prototype/shared_lib/request"
	"github.com/ma-zero-trust-prototype/shared_lib/response"
	"github.com/ma-zero-trust-prototype/shared_lib/structs"
	"net/http"
)

func CalcTrustScoresForAgent(agent *structs.NetworkAgent) {

	trustScores := sendRequestToTrustEngine(*agent)

	agent.User.TrustScore = trustScores.User.Score
	agent.Device.TrustScore = trustScores.Device.Score
	agent.TrustScore = trustScores.NetworkAgent.Score
}

func sendRequestToTrustEngine(agent structs.NetworkAgent) (trustScores structs.TrustScores) {

	client := &http.Client{}
	body := request.GetBodyFromStruct(agent)
	authorizationRequest, err := http.NewRequest(http.MethodPost, "https://"+env.GetTrustEngineAddress(), body)

	if err != nil {
		fmt.Println(err)
	}

	res, err := client.Do(authorizationRequest)
	response.ParseResponseBodyIntoStruct(res, &trustScores)

	return trustScores
}
