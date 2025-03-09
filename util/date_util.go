package util

import "time"

// FormatBirthDate formats a date to DD-MM-YYYY format.
func FormatBirthDate(date time.Time) string {
	return date.Format("02-01-2006")
}
