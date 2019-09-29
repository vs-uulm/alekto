package trustScore

import (
	"github.com/ma-zero-trust-prototype/shared_lib/stringUtil"
	"github.com/ma-zero-trust-prototype/shared_lib/structs"
	"github.com/ma-zero-trust-prototype/trust_engine/trustScore/deviceOpinion"
	"github.com/ma-zero-trust-prototype/trust_engine/trustScore/userOpinion"
)

func updateUserByAuthenticationAttempts(subject structs.Subject, variable structs.Variable, lastUpdated int64) structs.Variable {

	attempts := GetAuthenticationAttempts(subject, lastUpdated)

	if len(attempts) <= 0 {
		return variable
	}

	amount, failedAmount, ipAddresses := getMetadataOfAuthAttempts(attempts)

	userOpinion.UpdateByAuthAmount(amount, &variable)
	userOpinion.UpdateByFailedAuthAmount(failedAmount, &variable)
	userOpinion.UpdateByIpAddresses(len(ipAddresses), &variable)

	return variable
}

func updateDeviceByAuthenticationAttempts(subject structs.Subject, variable structs.Variable, lastUpdated int64) structs.Variable {

	attempts := GetAuthenticationAttempts(subject, lastUpdated)

	if len(attempts) <= 0 {
		return variable
	}

	amount, failedAmount, ipAddresses := getMetadataOfAuthAttempts(attempts)

	deviceOpinion.UpdateByDifferentUserAuthAmount(amount, &variable) // TODO
	deviceOpinion.UpdateByFailedAuthAmount(failedAmount, &variable)
	deviceOpinion.UpdateByIpAddresses(ipAddresses, &variable)
	deviceOpinion.UpdateByIpAddressAmount(len(ipAddresses), &variable)

	return variable
}

func getMetadataOfAuthAttempts(attempts []structs.AuthAttempt) (amount int, failedAmount int, ipAddressesData [][]string) {

	amount = len(attempts)
	ipAddressList := []string{}

	for _, attempt := range attempts {

		if attempt.Failed {
			failedAmount++
		}

		if !stringUtil.InSlice(attempt.IPAddress, ipAddressList) {
			ipAddressList = append(ipAddressList, attempt.IPAddress)

			ipAddressesData = append(ipAddressesData, []string{attempt.IPAddress, string(attempt.Timestamp)})
		}
	}

	return
}
