package userOpinion

import "github.com/ma-zero-trust-prototype/shared_lib/structs"

/**
 * Reduce Trustworthiness of an User if he authenticates more than 50 times in the last 24hours
 */
func UpdateByAuthAmount(amount int, variable *structs.Variable) {

	if amount > 10 {
		variable.Opinion = variable.Opinion.RaiseDisbelief(0.05)

	} else {
		variable.Opinion = variable.Opinion.RaiseBelief(0.05)
	}
}

/**
 * Microsoft recommends at least 4 attempts and no more than 10.
 * (https://docs.microsoft.com/en-us/previous-versions/windows/it-pro/windows-server-2012-R2-and-2012/hh994574(v=ws.11))
 */
func UpdateByFailedAuthAmount(failedAmount int, variable *structs.Variable) {

	maxAmount := 10
	minAmount := 4

	if failedAmount > maxAmount {
		variable.Opinion = variable.Opinion.RaiseDisbelief(0.1)

	} else if failedAmount > minAmount {
		variable.Opinion = variable.Opinion.RaiseDisbelief(0.05)

	} else if failedAmount <= minAmount {
		variable.Opinion = variable.Opinion.RaiseBelief(0.05)
	}
}

/**
 * Reduce Trustworthiness of an User if he authenticates from more than 4 ip addresses in the last 24hours
 */
func UpdateByIpAddresses(amount int, variable *structs.Variable) {

	if amount > 100 {
		variable.Opinion = variable.Opinion.RaiseDisbelief(0.5)

	} else if amount > 4 {
		variable.Opinion = variable.Opinion.RaiseDisbelief(0.1)

	} else if amount <= 4 {
		variable.Opinion = variable.Opinion.RaiseBelief(0.1)
	}
}
