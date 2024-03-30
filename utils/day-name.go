package utils

import (
	"strconv"
	"strings"
)

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

func ToDateIndo(date string) string {
	month := [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	// Split the date string
	dateSplit := strings.Split(date, "-")

	// Get the day, month, and year
	hari := dateSplit[2]
	monthInt, _ := strconv.Atoi(dateSplit[1])
	monthStr := month[monthInt-1]
	tahun := dateSplit[0]

	// Return the formatted date string
	return hari + " " + monthStr + " " + tahun
}

func TimeToIDDay(day string) string {
	switch day {
	case "Mon":
		return "Senin"
	case "Tue":
		return "Selasa"
	case "Wed":
		return "Rabu"
	case "Thu":
		return "Kamis"
	case "Fri":
		return "Jumat"
	case "Sat":
		return "Sabtu"
	case "Sun":
		return "Minggu"
	default:
		return "Tidak di ketahui"
	}
}
