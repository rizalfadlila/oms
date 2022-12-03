package util

import (
	"fmt"
	"time"
)

const (
	ISODateLayout = "2006-01-02"
	TimeLayout    = "15:04:05"
)

var (
	JakartaLocation = MustLoadLocation("Asia/Jakarta")
)

type ValidTime struct {
}

func MustLoadLocation(name string) *time.Location {
	loc, err := time.LoadLocation(name)
	if err != nil {
		panic(fmt.Errorf("failed to load location [%s] :%w", name, err))
	}

	return loc
}

func TimeInJakarta(t time.Time) time.Time {
	return t.In(JakartaLocation)
}

func DateNowInJakarta() string {
	return TimeInJakarta(time.Now()).Format(ISODateLayout)
}

func TimeNowInJakarta() string {
	return TimeInJakarta(time.Now()).Format(TimeLayout)
}

func DateTimeNowInJakarta() string {
	return TimeInJakarta(time.Now()).Format(fmt.Sprintf("%s %s", ISODateLayout, TimeLayout))
}

func (t ValidTime) Validate(value interface{}) error {
	date := fmt.Sprintf("2006-01-02T%v.00", value)
	_, err := time.Parse("2006-01-02T15:04:05", date)
	return err
}

func ValidationTime() ValidTime {
	return ValidTime{}
}
