package ip2location

import (
	"fmt"
)

/**
 * get Location by given Ip-Address string
 */
func GetLocationByIpAddress(ip string) Location {

	return getLocation(ip)
}

/**
 *	a := data.GetLocationByIpAddress("94.134.93.84")
 *  b := data.GetLocationByIpAddress("134.60.112.69")
 *	data.GetDistanceInKmOfTwoLocations(a, b)
 *
 * get the distance in km between two locations
 */
func GetDistanceInKmOfTwoLocations(locationA Location, locationB Location) int {

	a := getLatLongByLocation(locationA)
	b := getLatLongByLocation(locationB)

	distance := getDistanceInKmWithHaversineFormular(a, b)

	return distance
}

/**
 *	a := data.GetLocationByIpAddress("94.134.93.84")
 *  b := data.GetLocationByIpAddress("134.60.112.69")
 *	data.GetDistanceInKmOfTwoLocations(a, b)
 *
 * get the distance in km between two locations
 */
func GetDistanceOfTwoLatLongs(a LatLong, b LatLong) int {

	distance := getDistanceInKmWithHaversineFormular(a, b)

	fmt.Println(a, b, distance)

	return distance
}
