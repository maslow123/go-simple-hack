package jumlah_jam_kerja

import (
	"fmt"
	"strconv"
	"strings"
)

func JumlahJamKerja(time string) int {
	timeSplit := strings.Split(time, " ")
	if len(timeSplit) != 2 {
		fmt.Print("Invalid input\n")
		return 0
	}
	clockIn, err := strconv.Atoi(timeSplit[0])
	if err != nil || clockIn < 0 {
		fmt.Print("Invalid clock in\n")
		return 0
	}
	clockOut, err := strconv.Atoi(timeSplit[1])
	if err != nil || (clockOut > 24 || clockOut < clockIn) {
		fmt.Print("Invalid clock out\n")
		return 0
	}

	if clockIn > clockOut {
		fmt.Print("Clock in greather than clock out\n")
		return 0
	}

	overtime := getOvertime(clockIn, clockOut)
	hours := getWorkingHours(overtime, clockIn, clockOut)

	return hours
}

func getOvertime(clockInt, clockOut int) int {
	workingHours := clockOut - clockInt
	if workingHours > 8 {

		workingHours -= 8
		return workingHours
	}
	return 0
}

func getWorkingHours(overtime, clockIn, clockOut int) int {
	hours := clockOut - clockIn
	if overtime > 0 {
		hours = 8
		for i := 1; i <= overtime; i++ {
			if i < 3 {
				hours = hours + (1 * i)
			}
			if i >= 3 {
				hours = hours + (1 * 3)
			}
		}
	}

	return hours
}
