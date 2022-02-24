package utils

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	timeFormat2   = "2/1/2006"
	MaxDateString = "31/12/2999"
	MinDateString = "01/01/1900"
)

var UserLeapYear = false
var days2020 = []int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
var days2021 = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

// IsValidDate check date is valid
func IsValidDate(from, to string) int {
	//use regex to valid date
	re := regexp.MustCompile("(0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[012])/((19|20)\\d\\d)")
	if !re.MatchString(from) {
		fmt.Println(fmt.Sprintf("Use -h to help \nInvalid date from : %s\n", from))
		os.Exit(1)
	}
	if !re.MatchString(to) {
		fmt.Println(fmt.Sprintf("Use -h to help \nInvalid date from : %s\n", to))
		os.Exit(1)
	}

	//split date string to date
	fromSplit := strings.Split(from, "/")
	toSplit := strings.Split(to, "/")

	//get year from split date
	fromYear, _ := strconv.Atoi(fromSplit[2])
	toYear, _ := strconv.Atoi(toSplit[2])
	//check date is not more or less than we aspect
	if toYear > 2999 || toYear < 1900 {
		showErrorDateIsNotValid(toYear)
	}
	if fromYear > 2999 || fromYear < 1900 {
		showErrorDateIsNotValid(toYear)
	}
	//get month from split date
	fromMonth, _ := strconv.Atoi(fromSplit[1])
	toMonth, _ := strconv.Atoi(toSplit[1])

	//get day from split date
	fromDay, _ := strconv.Atoi(fromSplit[0])
	toDay, _ := strconv.Atoi(toSplit[0])

	//total day of year in between tow different years
	//check we have tow different year
	isFromYear, dayOfYear := calculateAllDayOFYear(fromYear, toYear)

	//get all day from between different month
	dayOfMonthTotal := 0

	//if we dont have tow different years
	// we just calculated days between different month
	if !isFromYear {
		dayOfMonthTotal = calculateDayOfMonth(fromYear, toYear, fromMonth, toMonth, dayOfMonthTotal, toDay, fromDay)
	} else {
		//if we have tow different month we must find number of days we added extra to dayOfYear
		//then we subtract extra day from day of year
		if toMonth != 1 {
			dd := calculateDifferenceDay(toYear, toYear, toMonth, 12, dayOfMonthTotal, toDay, 0)
			dayOfYear -= dd
		}
		if fromMonth != 1 {
			dd2 := calculateDayOfMonth(fromYear, fromYear, 1, fromMonth, dayOfMonthTotal, 0, fromDay)
			dayOfYear -= dd2
		}

		//check if we have ont the first mnont
		//subtract fromDay from dayOfYear
		if fromMonth == 1 {
			dayOfYear -= fromDay
		}
		//check if toMonth is equle to 1 and toDay grater than 1
		//we get month day base on leapYear year and subtracted from day of year
		if toMonth == 1 && toDay != 1 {
			if leapYear(toYear) {
				val := 0
				val = days2020[toMonth-1] - toDay
				dayOfYear -= val
			} else {
				val := 0
				val = days2021[toMonth-1] - toDay
				dayOfYear -= val

			}
		}
	}
	//subtract daffy wight dayOfMonthTotal
	sum := dayOfYear + dayOfMonthTotal
	if sum < 0 {
		sum *= -1
	}
	// subtract with 1
	sum -= 1
	return sum

}

//show error and exit
func showErrorDateIsNotValid(toYear int) {
	fmt.Println(fmt.Sprintf("Use -h to help \n Date is hiegher than date we expect : %s\n", toYear))
	os.Exit(1)
}

func calculateAllDayOFYear(fromYear int, toYear int) (bool, int) {
	allDayOfYear := 0
	var isFromYear = false
	//for find calculate days between to different year we we start from bigger year
	y := fromYear
	if fromYear < toYear {
		fromYear = toYear
		toYear = y
	}
	if fromYear > toYear {
		isFromYear = true
		for i := toYear; i < fromYear; i++ {
			//check is leap year
			if leapYear(i) {
				for _, days := range days2020 {
					//add all of day from array of month
					allDayOfYear += days
				}
			} else {
				for _, days := range days2021 {
					//add all of day from array of month
					allDayOfYear += days
				}
			}

		}
	}

	return isFromYear, allDayOfYear
}

//calculate all days between months
func calculateDayOfMonth(fromYear int, toYear int, fM int, tm int, dayOfYMonth int, toDate int, fromDate int) int {
	//check is leap year and calculate all days between months
	if leapYear(fromYear) && leapYear(toYear) {
		lastIndex := 0
		for i := fM; i <= tm; i++ {
			lastIndex = i
			dayOfYMonth += days2020[(i - 1)]
		}
		d := days2020[lastIndex-1] - toDate
		dayOfYMonth -= d
		dayOfYMonth -= fromDate
	} else {
		lastIndex := 0
		for i := fM; i <= tm; i++ {
			lastIndex = i
			dayOfYMonth += days2021[(i - 1)]
		}
		d := days2021[lastIndex-1] - toDate
		dayOfYMonth -= d
		dayOfYMonth -= fromDate
	}
	return dayOfYMonth
}
//calculate all days between days
func calculateDifferenceDay(fromYear int, toYear int, fromMonth int, toMonth int, dayOfYMonth int, toDate int, fromDate int) int {
	//check is leap year and calculate all days between days
	if leapYear(fromYear) && leapYear(toYear) {
		for i := fromMonth; i <= toMonth; i++ {
			dayOfYMonth += days2020[(i - 1)]
		}
		dayOfYMonth -= toDate

		dayOfYMonth -= fromDate
	} else {

		for i := fromMonth; i <= toMonth; i++ {

			dayOfYMonth += days2021[(i - 1)]
		}
		dayOfYMonth -= toDate
		dayOfYMonth -= fromDate
	}
	return dayOfYMonth
}

// Check is leap year
func leapYear(year int) bool {
	//if this leap year is active is setting another result
	if UserLeapYear {
		return ((year%4 == 0) && (year%100 != 0)) || (year%400 == 0)
	}
	return false
}
