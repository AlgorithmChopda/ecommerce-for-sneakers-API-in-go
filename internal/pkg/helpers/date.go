package helpers

import "time"

func ParseDate(dateString string) (time.Time, error) {
	parsedDate, err := time.Parse("02-01-2006", dateString)
	if err != nil {
		return time.Time{}, err
	}
	return parsedDate, nil
}
