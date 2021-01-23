package util

import (
	"time"
)

//ConvertStringToDate it converts the date to time
// making this function to available across the app so that if we have any change
// in the dateformat it is easily changeable accross app.
func ConvertStringToDate(dateString string) (time.Time, error) {
	return time.Parse("2006-01-02T15:04", dateString)
}

//GetLastDayOfMonth calculates to return the last day of the given time-date
func GetLastDayOfMonth(t time.Time) int {
	return int(time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local).AddDate(0, 1, -1).Day())
}
