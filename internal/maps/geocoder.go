package maps

import (
	"context"
	"log"
	"os"

	"googlemaps.github.io/maps"
)

type Address struct {
	AddressComponents []AddressComponent `json:"addressComponents"`
	FormattedAddress  string             `json:"formattedAddress"`
	Latitude          float64            `json:"latitude"`
	Longitude         float64            `json:"longitude"`
}

type AddressComponent struct {
	LongName  string   `json:"long_name"`
	ShortName string   `json:"short_name"`
	Types     []string `json:"types"`
}

func SearchPlace(place string) (Address, string) {
	opts := maps.WithAPIKey(os.Getenv("GOOGLE_MAPS_API_KEY"))
	client, err := maps.NewClient(opts)

	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	request := &maps.GeocodingRequest{Address: place}
	response, err := client.Geocode(context.Background(), request)

	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	formatted := formatResponse(response[0])

	return formatted, formatted.FormattedAddress
}

func formatResponse(response maps.GeocodingResult) Address {
	addrComponents := []AddressComponent{}

	for _, component := range response.AddressComponents {
		addrComponents = append(addrComponents, AddressComponent(component))
	}

	return Address{
		AddressComponents: addrComponents,
		FormattedAddress:  response.FormattedAddress,
		Latitude:          response.Geometry.Location.Lat,
		Longitude:         response.Geometry.Location.Lng,
	}

}
