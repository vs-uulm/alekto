package trustScore

import (
	"fmt"
	"github.com/ma-zero-trust-prototype/shared_lib/data/ip2location"
	"github.com/ma-zero-trust-prototype/shared_lib/structs"
	"testing"
)

func TestCalcTrustScoresForAgent(t *testing.T) {

	var agent = structs.NetworkAgent{
		User: structs.User{
			Id:             "bender",
			Role:           "student",
			Category:       "standard",
			Authentication: "basicAuth"},
		Device: structs.Device{
			Id:             "bendersDevice",
			Type:           "laptop",
			Authentication: "mtls",
			OS:             "Ubuntu",
			Owner:          "bender",
			IPAddress:      "134.60.112.69",
			Location:       ip2location.GetLocationByIpAddress("134.60.112.69")}}

	fmt.Printf("Resulting Agent: %+v \n", agent)

	scores := CalcTrustScoresForAgent(agent)

	fmt.Printf("Resulting scores: %+v \n", scores)
}
