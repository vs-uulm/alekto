package deviceOpinion

import (
	"github.com/ma-zero-trust-prototype/shared_lib/data/ip2location"
	"github.com/ma-zero-trust-prototype/shared_lib/sliceUtil"
	"github.com/ma-zero-trust-prototype/shared_lib/stringUtil"
	"github.com/ma-zero-trust-prototype/shared_lib/structs"
)


/**
 * Reduce Trustworthiness of an Device if there are more than
 * 50 different user authentications
 * TODO in 24h
 */
func UpdateByDifferentUserAuthAmount (amount int, variable *structs.Variable) {

	if amount > 10 {

		variable.Opinion = variable.Opinion.RaiseDisbelief(0.05)
	}
}

/**
 * raise the disbelief of an device, if too many failed authentication attempts happened
 */
func UpdateByFailedAuthAmount(amount int, variable *structs.Variable) {

	if amount > 100 {
		variable.Opinion = variable.Opinion.RaiseDisbelief(0.2)
	} else if amount > 10 {
		variable.Opinion = variable.Opinion.RaiseDisbelief(0.1)
	} else if amount < 3 {
		variable.Opinion = variable.Opinion.RaiseBelief(0.05)
	}
}

/**
 * raise the belief or disbelief of an device by it's travel speed (location changes)
 */
func UpdateByIpAddresses(ipAddresses [][]string, variable *structs.Variable) {

	suspiciousIpChanges := 0

	for index := len(ipAddresses) - 1; index >= 0; index-- {

		address := ipAddresses[index]
		location := ip2location.GetLocationByIpAddress(address[0])
		time := address[1]

		suspiciousIpChanges += getSuspiciousMovements(location, time, ipAddresses)

		if len(ipAddresses) > 1 {
			ipAddresses = deleteObservedElement(ipAddresses, index)
		} else {
			ipAddresses = [][]string{}
		}
	}

	if suspiciousIpChanges > 10 {
		variable.Opinion = variable.Opinion.RaiseDisbelief(1)
	} else if suspiciousIpChanges > 0 {
		variable.Opinion = variable.Opinion.RaiseDisbelief(0.5)
	} else if suspiciousIpChanges == 0 {
		variable.Opinion = variable.Opinion.RaiseBelief(0.05)
	}
}


/**
 * raise the disbelief of an device by it's ip addresses
 * TODO Pro 24h
 */
func UpdateByIpAddressAmount(amount int, variable *structs.Variable) {

	if amount > 100 {
		variable.Opinion = variable.Opinion.RaiseDisbelief(0.2)
	} else if amount > 10 {
		variable.Opinion = variable.Opinion.RaiseDisbelief(0.1)
	} else if amount < 3 {
		variable.Opinion = variable.Opinion.RaiseBelief(0.05)
	}
}

/**
 * check if a device moves with more than 1000 km/h
 */
func getSuspiciousMovements (location ip2location.Location, time string, addresses [][]string) int {
	suspiciousMovements := 0
	maximalRealisticTravelSpeed := 1000

	for _, other := range addresses {

		otherLocation := ip2location.GetLocationByIpAddress(other[0])

		distanceInKm := ip2location.GetDistanceInKmOfTwoLocations(location, otherLocation)
		timeDistanceInS := stringUtil.ToInt(time) - stringUtil.ToInt(other[1])
		speed := calcKmPerHour(distanceInKm, timeDistanceInS)

		if speed > maximalRealisticTravelSpeed {
			suspiciousMovements++
		}
	}

	return suspiciousMovements
}

/*
 * calculate km/h from distance in km and time in seconds
 */
func calcKmPerHour(distanceInKm int, timeInS int) int {

	if timeInS < 0 {
		timeInS *= -1
	}

	timeInH := timeInS / (60 * 60)

	if timeInH == 0 {
		timeInH = 1
	}

	speed := distanceInKm / timeInH

	return speed
}

func deleteObservedElement (ipAddresses [][]string, index int) [][]string {

	return sliceUtil.DropElementAtIndex(ipAddresses, index)
}