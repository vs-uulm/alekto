package ip2location

import (
	"flag"
	"github.com/ip2location/ip2location-go"
)

const (
	r = 6378100 // Earth radius in METERS
)

type Location struct {
	CountryCode string  `yaml:"countrycode"`
	Region      string  `yaml:"region"`
	City        string  `yaml:"city"`
	Lat         float32 `yaml:"lat"`
	Long        float32 `yaml:"long"`
	ZipCode     string  `yaml:"zipcode"`
	Radius      int     `yaml:"radius"`
}

func getDbPath() string {

	dbPath := "./data/ip2location/bin/IP2LOCATION-LITE-DB9.BIN"

	if flag.Lookup("test.v") != nil {
		dbPath = "../data/ip2location/bin/IP2LOCATION-LITE-DB9.BIN"
	}

	return dbPath
}

/**
 * get Location by given Ip-Address string
 */
func getLocation(ip string) (location Location) {

	ip2location.Open(getDbPath())

	defer ip2location.Close()

	results := ip2location.Get_all(ip)

	location = Location{
		CountryCode: results.Country_short,
		Region:      results.Region,
		City:        results.City,
		Lat:         results.Latitude,
		Long:        results.Longitude,
		ZipCode:     results.Zipcode}

	return
}
