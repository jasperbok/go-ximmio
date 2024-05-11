package ximmio

type Calendar struct {
	PickupDates    []Time `json:"pickupDates"`
	PickupType     int    `json:"pickupType"`
	PickupTypeText string `json:"_pickupTypeText"`
}

type Calendars struct {
	Calendars []Calendar `json:"dataList"`
}
