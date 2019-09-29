package validator

import (
	"github.com/ma-zero-trust-prototype/shared_lib/data/ip2location"
)

func validateLocation(location ip2location.Location, userLocation ip2location.Location) (valid bool) {

	if location.Lat != 0 && location.Long != 0 {
		return validateLatLong(location, userLocation)
	}

	if location.CountryCode != "" || location.Region != "" || location.City != "" {
		return validateAddress(location, userLocation)
	}

	return false
}

func validateLatLong(location ip2location.Location, userLocation ip2location.Location) bool {

	requiredRadiusInKm := location.Radius

	if requiredRadiusInKm <= 0 {
		requiredRadiusInKm = 50
	}

	distanceInKm := ip2location.GetDistanceInKmOfTwoLocations(location, userLocation)

	return distanceInKm <= requiredRadiusInKm
}

/**
 * TODO Get Lat Long of address
 */
func validateAddress(location ip2location.Location, userLocation ip2location.Location) (valid bool) {

	valid = true

	if location.CountryCode != "" {
		valid = valid && location.CountryCode == userLocation.CountryCode
	}

	if location.City != "" {
		valid = valid && location.City == userLocation.City
	}

	if location.Region != "" {
		valid = valid && location.Region == userLocation.Region
	}

	return valid
}
