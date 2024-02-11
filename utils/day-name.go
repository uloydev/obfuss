package utils

func GetDayName(str string) string {
	var result string = "Monday"

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

	return result
}
