package fields

type Kind = string

const (
	UserAuthentication     Kind = "userAuthentication"
	User                   Kind = "user"
	UserRole               Kind = "userRole"
	UserCategory           Kind = "userCategory"
	UserTrustScore         Kind = "userTrustscore"
	Device                 Kind = "device"
	DeviceAuthentication   Kind = "deviceAuthentication"
	DeviceLocation         Kind = "deviceLocation"
	DeviceTrustScore       Kind = "deviceTrustscore"
	DeviceType             Kind = "deviceType"
	NetworkAgent           Kind = "networkagent"
	NetworkAgentTrustScore Kind = "networkagentTrustscore"
	Policy                 Kind = "policy"
)
