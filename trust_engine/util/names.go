package util

import (
	"github.com/ma-zero-trust-prototype/shared_lib/enum/fields"
	"github.com/ma-zero-trust-prototype/shared_lib/structs"
	"strings"
)

func GetNetworkAgentsName(agent structs.NetworkAgent) string {
	return agent.User.Id + "_" + agent.Device.Id
}

func GetUnIdentifiedDeviceName(agent structs.NetworkAgent) string {
	return agent.User.Id + "s_UnIdentifiedDevice"
}

func IsUnIdentifiedDevice(subject structs.Subject) bool {
	return subject.Kind == fields.Device &&
		strings.Contains(subject.Name, "s_UnIdentifiedDevice")
}
