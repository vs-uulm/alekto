package ip2location

import "math"

type LatLong struct {
	Lat  float64 `yaml:"lat"`
	Long float64 `yaml:"long"`
}

/**
 * get lat long struct by given location
 */
func getLatLongByLocation(location Location) LatLong {

	var lat, long float64

	lat = float64(location.Lat) * math.Pi / 180
	long = float64(location.Long) * math.Pi / 180

	return LatLong{lat, long}
}

/**
 * http://en.wikipedia.org/wiki/Haversine_formula
 * calc h sin
 */
func hSin(theta float64) float64 {

	return math.Pow(math.Sin(theta/2), 2)
}

func getDistanceInKmWithHaversineFormular(a LatLong, b LatLong) int {

	haversine := math.Cos(a.Lat)*math.Cos(b.Lat)*hSin(b.Long-a.Long) + hSin(b.Lat-a.Lat)
	d := 2 * r * math.Asin(math.Sqrt(haversine)) / 1000

	return int(d)
}
