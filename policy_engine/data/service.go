package data

import (
	"github.com/ma-zero-trust-prototype/shared_lib/data/ip2location"
	"github.com/ma-zero-trust-prototype/shared_lib/structs"
)

func GetUserAndDeviceInfoForAgent(agent *structs.NetworkAgent) {

	dummyAgent := getUserDummyData(agent.User.Id)

	agent.User.Role = dummyAgent.User.Role
	agent.User.Category = dummyAgent.User.Category
	agent.Device.Type = dummyAgent.Device.Type
	agent.Device.OS = "Ubuntu"
	agent.Device.Owner = "leela"
	agent.Device.Location = GetLocationByIpAddress(agent.Device.IPAddress)
}

/**
 * get Location by given Ip-Address string
 */
func GetLocationByIpAddress(ip string) ip2location.Location {

	return ip2location.GetLocationByIpAddress(ip)
}

func getUserDummyData(user string) structs.NetworkAgent {

	users := make(map[string]structs.NetworkAgent)

	users["default"] = structs.NetworkAgent{
		User: structs.User{
			Role:     "student",
			Category: "lowPrivileged"},
		Device: structs.Device{
			Type: "stationary"}}

	users["bender"] = structs.NetworkAgent{
		User: structs.User{
			Role:     "student",
			Category: "lowPrivileged"},
		Device: structs.Device{
			Type: "stationary"}}

	users["leela"] = structs.NetworkAgent{
		User: structs.User{
			Role:     "student",
			Category: "lowPrivileged"},
		Device: structs.Device{
			Type: "mobile"}}

	users["hermes"] = structs.NetworkAgent{
		User: structs.User{
			Role:     "student",
			Category: "lowPrivileged"},
		Device: structs.Device{
			Type: "stationary"}}

	users["professor"] = structs.NetworkAgent{
		User: structs.User{
			Role:     "teacher",
			Category: "privileged"},
		Device: structs.Device{
			Type: "stationary"}}

	users["zoidberg"] = structs.NetworkAgent{
		User: structs.User{
			Role:     "admin",
			Category: "highPrivileged"},
		Device: structs.Device{
			Type: "stationary"}}

	if val, ok := users[user]; ok {
		return val
	} else {
		return users["default"]
	}
}
