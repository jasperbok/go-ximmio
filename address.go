package ximmio

// Address is a collection of address information as returned by Ximmio's
// Waste API.
//
// The actual API response also has a field `BuildingCategory`. However, at the
// time of writing this client has only been tested with residential addresses,
// for which the field always seem to be `null`. So no data type can be defined
// for it yet.
type Address struct {
	ID                    int    `json:"ID"`
	UniqueID              string `json:"UniqueId"`
	Street                string `json:"Street"`
	HouseNumber           string `json:"HouseNumber"`
	HouseLetter           string `json:"HouseLetter"`
	HouseNumberIndication string `json:"HouseNumberIndication"`
	HouseNumberAddition   string `json:"HouseNumberAddition"`
	PostalCode            string `json:"ZipCode"`
	City                  string `json:"City"`
	Community             string `json:"Community"`
}
