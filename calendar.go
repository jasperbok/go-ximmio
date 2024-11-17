package ximmio

type Calendar struct {
	PickupDates    []Time `json:"pickupDates"`
	PickupType     int    `json:"pickupType"`
	PickupTypeText string `json:"_pickupTypeText"`
	Description    string `json:"description"`
}

type Calendars struct {
	Calendars []Calendar `json:"dataList"`
}
