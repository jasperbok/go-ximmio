package ximmio

// Address is a collection of address information as returned by Ximmio's
// Waste API.
type Address struct {
	UniqueID    string `json:"UniqueId"`
	Street      string `json:"Street"`
	HouseNumber string `json:"HouseNumber"`
	HouseLetter string `json:"HouseLetter"`
	ZipCode     string `json:"ZipCode"`
	City        string `json:"City"`
	Community   string `json:"Community"`
}
