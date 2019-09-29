package deviceRequirements

import (
	"github.com/ma-zero-trust-prototype/policy_engine/policies/validator"
	"github.com/ma-zero-trust-prototype/shared_lib/enum/fields"
	"github.com/ma-zero-trust-prototype/shared_lib/structs"
	"github.com/ma-zero-trust-prototype/shared_lib/structs/policy"
)

func Validate(device policy.Device, agent structs.Device) (success bool, message string) {

	success, message = validator.ValidateStringArrayBySingleValue(fields.DeviceType, device.Type, agent.Type)
	if !success {
		return
	}

	success, message = validator.ValidateStringArrayBySingleValue(fields.DeviceAuthentication, device.Authentication, agent.Authentication)
	if !success {
		return
	}

	success, message = validator.ValidateLocation(fields.DeviceLocation, device.Location, agent.Location)
	if !success {
		return
	}

	success, message = validator.ValidateTrustScore(fields.DeviceTrustScore, device.TrustScore, agent.TrustScore)

	return
}
