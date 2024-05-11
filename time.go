package ximmio

import (
	"strings"
	"time"
)

const (
	XIMMIO_TIME_LAYOUT = "2006-01-02T15:04:05"
)

type Time time.Time

func (t Time) MarshalJSON() ([]byte, error) {
	formatted := time.Time(t).Format(XIMMIO_TIME_LAYOUT)
	return []byte(formatted), nil
}

func (t *Time) UnmarshalJSON(b []byte) error {
	_t, err := time.Parse(XIMMIO_TIME_LAYOUT, strings.Trim(string(b), "\""))
	if err != nil {
		return err
	}

	*t = Time(_t)

	return nil
}
