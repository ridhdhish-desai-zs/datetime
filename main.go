package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	dailyFunc "github.com/ridhdhish-desai-zs/datetime/Daily"
	"github.com/ridhdhish-desai-zs/datetime/Monthly"
	"github.com/ridhdhish-desai-zs/datetime/Weekly"
	"github.com/ridhdhish-desai-zs/datetime/Yearly"
)

// Returns days/dates/months with [interval]
func getParams(val string) (*[]int, *int) {
	params := strings.Split(val, "/")

	var daysOfTheWeek []int
	var interval int

	week := params[0]

	if week != "*" {
		if week[:1] == "(" {
			between := strings.Split(week[1:len(week)-1], "-")
			from, _ := strconv.Atoi(between[0])
			to, _ := strconv.Atoi(between[1])

			for i := from; i <= to; i++ {
				daysOfTheWeek = append(daysOfTheWeek, i)
			}
		} else {
			days := strings.Split(week, ",")

			for _, v := range days {
				val, _ := strconv.Atoi(v)
				daysOfTheWeek = append(daysOfTheWeek, val)
			}
		}

		if len(params) > 1 {
			interval, _ = strconv.Atoi(params[1])
			return &daysOfTheWeek, &interval
		} else {
			return &daysOfTheWeek, nil
		}
	} else {
		if len(params) > 1 {
			interval, _ = strconv.Atoi(params[1])
			return nil, &interval
		} else {
			return nil, nil
		}
	}
}

func checkRecurrenceType(weekDays, dates, daily, months *[]int, weekInterval, datesInterval, dailyInterval, monthInterval *int) string {
	// Weekly occurrence
	if weekDays != nil || weekInterval != nil {
		if dates != nil || months != nil {
			return "Invalid"
		}
		return "weekly"
		// Daily occurrence
	} else if datesInterval != nil && dates == nil {
		if weekDays != nil || months != nil {
			return "Invalid"
		}
		return "daily"
		// Monthly occurrence
	} else if dates != nil && monthInterval != nil && months == nil {
		if weekDays != nil {
			return "Invalid"
		}
		return "monthly"
		// Yearly occurrence
	} else if dates != nil && monthInterval != nil && months != nil {
		if weekDays != nil {
			return "Invalid"
		}
		return "yearly"
	}

	return "Invalid"
}

func main() {
	// regex := "00 16 1,2,6/2 * *"
	// regex := "00 16 * 1 */10"
	regex := "00 00 * 1 4/120"
	// regex := "00 16 * */10 *"

	regexParams := strings.Split(regex, " ")

	weekDays, weekInterval := getParams(regexParams[2])
	daily, dailyInterval := getParams(regexParams[3])
	dates, datesInterval := getParams(regexParams[3])
	months, monthInterval := getParams(regexParams[4])

	hours, _ := strconv.Atoi(regexParams[1])
	minutes, _ := strconv.Atoi(regexParams[0])

	endDate, _ := time.Parse(time.RFC3339, "2022-04-23T10:00:00Z")
	startDate, _ := time.Parse(time.RFC3339, "2022-02-28T15:00:00Z")
	currentDate, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	if currentDate.Sub(startDate) > 0 {
		startDate = currentDate
	}

	recurrenceType := checkRecurrenceType(weekDays, dates, daily, months, weekInterval, datesInterval, dailyInterval, monthInterval)
	fmt.Println(recurrenceType)

	switch recurrenceType {
	case "daily":
		nextDues := dailyFunc.GetDailyDues(hours, minutes, *dailyInterval, startDate, endDate)
		fmt.Println(nextDues)
	case "weekly":
		nextDues := Weekly.GetFirstDueDate(hours, minutes, *weekInterval, startDate, endDate, *weekDays)
		finalDues := Weekly.GetNextDues(*weekInterval, nextDues, endDate)

		fmt.Println(finalDues)
	case "monthly":
		nextDues := Monthly.GetFirstDateOccurrence(hours, minutes, *monthInterval, startDate, endDate, *dates)
		finalDues := Monthly.GetNextDues(*monthInterval, endDate, nextDues)

		fmt.Println(finalDues)
	case "yearly":
		nextDues := Yearly.GetFirstDateOccurrence(hours, minutes, *monthInterval, startDate, endDate, *dates, *months)
		finalDues := Yearly.GetNextDues(*monthInterval, endDate, nextDues)

		fmt.Println(finalDues)
	}
}
