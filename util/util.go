package util

import (
	"fmt"
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

//CheckPrime checks whether the given number is prime or not
func CheckPrime(number int) bool {
	isPrime := true
	if number == 0 || number == 1 {
		fmt.Printf(" %d is not a  prime number\n", number)
		return false
	}
	for i := 2; i <= number/2; i++ {
		if number%i == 0 {
			fmt.Printf(" %d is not a  prime number\n", number)
			isPrime = false
			break
		}
	}
	return isPrime
}
