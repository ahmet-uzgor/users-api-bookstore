package dateutil

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
	apiDbLayout   = "2006-01-02 15:04:05"
)

func GetTodayDate() time.Time {
	return time.Now().UTC()
}

func GetTodayString() string {
	return GetTodayDate().Format(apiDateLayout)
}

func GetTodayDbString() string {
	return GetTodayDate().Format(apiDbLayout)
}
