package device

type deviceType string

const (
	Default    deviceType = "all"
	Mobile     deviceType = "mobile"
	Stationary deviceType = "stationary"
	Managed    deviceType = "managed"
)
