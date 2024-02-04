package utils

import "errors"

func GetDayName(str string) (string, error) {
	var (
		errorFailedParse = errors.New("failed to parse day")
		result           string
	)

	switch str {
	case "Senin":
		result = "Monday"

	case "Selasa":
		result = "Tuesday"

	case "Rabu":
		result = "Wednesday"

	case "Kamis":
		result = "Thursday"

	case "Jumat":
		result = "Friday"

	case "Sabtu":
		result = "Saturday"

	case "Minggu":
		result = "Sunday"
	}

	return result, errorFailedParse
}
