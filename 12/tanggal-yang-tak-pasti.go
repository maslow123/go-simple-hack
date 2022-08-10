package tanggal_yang_tak_pasti

import (
	"fmt"
	"strconv"
	"strings"
)

func TanggalYangTakPasti(date string) string {
	// Format date only accept DD-MM-YYYY, DD MM YYYY, DD/MM/YYYY
	separators := []string{"-", " ", "/"}

	// check format date is valid or not
	formatDateIsValid, dateSplit := formatDateIsValid(separators, date)
	if !formatDateIsValid {
		fmt.Print("Kesalahan format tanggal\n")
		return "TIDAK"
	}
	// check date, month, and year is valid or not
	m, y, filterDate := getFilterDate(dateSplit)

	// if m is empty, find date and month
	if m == "" {
		month := filterDate[0]
		dateInt := filterDate[1]
		// check if the month and dateInt <= 12, return TIDAK
		if month <= 12 && dateInt <= 12 {
			fmt.Printf("Ada dua kemungkinan..\n")
			return "TIDAK"
		}

		// check the values <= 12 that is the month.
		if month >= 12 && dateInt < month {
			month = dateInt
			dateInt = month
		}

		fmt.Printf("Date: %d Month: %d Year: %s\n", dateInt, month, y)
	}

	return "YA"
}

func monthList() map[string]string {
	month := map[string]string{
		"jan": "jan",
		"feb": "feb",
		"mar": "mar",
		"apr": "apr",
		"mei": "may",
		"jun": "jun",
		"jul": "jul",
		"agu": "aug",
		"sep": "sep",
		"okt": "oct",
		"nov": "nov",
		"des": "dec",
	}

	return month
}

func getFilterDate(dateSplit []string) (string, string, []int) {
	var m, y string
	var filterDate []int
	// iterate the dateSplit
	for _, str := range dateSplit {
		// can parsing to string or not? if cant that is the month.
		intV, err := strconv.Atoi(str)
		if err != nil {
			m = str
			continue
		}

		// check str is year or not
		if err == nil && len(str) == 4 {
			y = str
			continue
		}
		filterDate = append(filterDate, intV)
	}

	return m, y, filterDate
}

func formatDateIsValid(separators []string, date string) (bool, []string) {
	formatDateIsValid := false
	var dateSplit []string
	for _, separator := range separators {
		dateSplit = strings.Split(date, separator)
		if len(dateSplit) == 3 {
			formatDateIsValid = true
			break
		}
	}

	return formatDateIsValid, dateSplit
}
