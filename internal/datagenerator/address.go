package datagenerator

import (
	"math/rand"
	"time"
)

type Address struct {
	AddressComponents []AddressComponents `json:"addressComponents"`
	FormattedAddress  string              `json:"formattedAddress"`
	Latitude          float32             `json:"latitude"`
	Longitude         float32             `json:"longitude"`
}

type AddressComponents struct {
	LongName  string   `json:"long_name"`
	ShortName string   `json:"short_name"`
	Types     []string `json:"types"`
}

func GetRandomAddress() (Address, string) {
	rand.Seed(time.Now().UnixNano())

	components := []Address{
		{
			AddressComponents: getMtlAddressComponents(),
			FormattedAddress:  "Montreal, QC, Canada",
			Latitude:          45.5018869,
			Longitude:         -73.56739189999999,
		},
		{
			AddressComponents: getLaVouteAddress(),
			FormattedAddress:  "Montréal, QC, Canada",
			Latitude:          45.5022567,
			Longitude:         -73.559135,
		},
		{
			AddressComponents: getParisAddressComponents(),
			FormattedAddress:  "Paris, France",
			Latitude:          48.856614,
			Longitude:         2.3522219,
		},
	}

	max := len(components) - 1

	addressComponent := components[rand.Intn(max+1)]

	return addressComponent, addressComponent.FormattedAddress
}

func getMtlAddressComponents() []AddressComponents {
	return []AddressComponents{
		{
			LongName:  "Montreal",
			ShortName: "Montreal",
			Types: []string{
				"locality",
				"political",
			},
		},
		{
			LongName:  "Montréal",
			ShortName: "Montréal",
			Types: []string{
				"administrative_area_level_3",
				"political",
			},
		},
		{
			LongName:  "Montreal",
			ShortName: "Montreal",
			Types: []string{
				"administrative_area_level_2",
				"political",
			},
		},
		{
			LongName:  "Quebec",
			ShortName: "QC",
			Types: []string{
				"administrative_area_level_1",
				"political",
			},
		},
		{
			LongName:  "Canada",
			ShortName: "CA",
			Types: []string{
				"country",
				"political",
			},
		},
	}
}

func getLaVouteAddress() []AddressComponents {
	return []AddressComponents{
		{
			LongName:  "Montréal",
			ShortName: "Montréal",
			Types: []string{
				"locality",
				"political",
			},
		},
		{
			LongName:  "Ville-Marie",
			ShortName: "Ville-Marie",
			Types: []string{
				"political",
				"sublocality",
				"sublocality_level_1",
			},
		},
		{
			LongName:  "Montréal",
			ShortName: "Montréal",
			Types: []string{
				"administrative_area_level_3",
				"political",
			},
		},
		{
			LongName:  "Communauté-Urbaine-de-Montréal",
			ShortName: "Communauté-Urbaine-de-Montréal",
			Types: []string{
				"administrative_area_level_2",
				"political",
			},
		},
		{
			LongName:  "Québec",
			ShortName: "QC",
			Types: []string{
				"administrative_area_level_1",
				"political",
			},
		},
		{
			LongName:  "Canada",
			ShortName: "CA",
			Types: []string{
				"country",
				"political",
			},
		},
		{
			LongName:  "H2Y 1P5",
			ShortName: "H2Y 1P5",
			Types: []string{
				"postal_code",
			},
		},
	}
}

func getParisAddressComponents() []AddressComponents {
	return []AddressComponents{
		{
			LongName:  "Paris",
			ShortName: "Paris",
			Types: []string{
				"locality",
				"political",
			},
		},
		{
			LongName:  "Paris",
			ShortName: "Paris",
			Types: []string{
				"administrative_area_level_2",
				"political",
			},
		},
		{
			LongName:  "Île-de-France",
			ShortName: "IDF",
			Types: []string{
				"administrative_area_level_1",
				"political",
			},
		},
		{
			LongName:  "France",
			ShortName: "FR",
			Types: []string{
				"country",
				"political",
			},
		},
	}
}
