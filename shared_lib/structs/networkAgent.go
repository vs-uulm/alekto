package structs

import (
	"github.com/ma-zero-trust-prototype/shared_lib/data/ip2location"
	"github.com/ma-zero-trust-prototype/shared_lib/request"
)

type NetworkAgent struct {
	User       User
	Device     Device
	TrustScore TrustScore
}

type User struct {
	Id             string
	Role           string
	Category       string
	Authentication string
	TrustScore     TrustScore
}

type Device struct {
	Id             string
	Type           string
	Authentication string
	OS             string
	IPAddress      string
	Location       ip2location.Location
	Owner          string
	TrustScore     TrustScore
}

/**
 * Create an User NetworkAgent from authorization payload
 */
func CreateAgentFromAuthorizationPayload(payloadStruct request.AuthorizationPayload) (agent NetworkAgent) {

	agent = NetworkAgent{
		User: User{
			Id:             payloadStruct.Username,
			Authentication: payloadStruct.UserAuthentication,
		},
		Device: Device{
			Id:             payloadStruct.DeviceId,
			IPAddress:      payloadStruct.IPAddress,
			Authentication: payloadStruct.DeviceAuthentication}}

	return
}
