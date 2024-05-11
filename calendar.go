package ximmio

import "time"

type Calendar struct {
	PickupDates    []time.Time `json:"pickupDates"`
	PickupType     int         `json:"pickupType"`
	PickupTypeText string      `json:"_pickupTypeText"`
}

type Calendars struct {
	Calendars []Calendar `json:"dataList"`
}
