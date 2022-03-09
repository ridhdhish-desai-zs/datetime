package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	dailyFunc "github.com/ridhdhish-desai-zs/datetime/Daily"
	"github.com/ridhdhish-desai-zs/datetime/Monthly"
	"github.com/ridhdhish-desai-zs/datetime/Weekly"
	"github.com/ridhdhish-desai-zs/datetime/Yearly"
)

// Returns days/dates/months with [interval]
func getParams(input string) ([]int, *int) {
	params := strings.Split(input, "/")

	var values []int // stores days/dates/months
	var interval int

	inputValues := params[0]

	if inputValues != "*" {
		if inputValues[:1] == "(" {
			between := strings.Split(inputValues[1:len(inputValues)-1], "-")
			from, _ := strconv.Atoi(between[0])
			to, _ := strconv.Atoi(between[1])

			for i := from; i <= to; i++ {
				values = append(values, i)
			}
		} else {
			days := strings.Split(inputValues, ",")

			for _, v := range days {
				val, _ := strconv.Atoi(v)
				values = append(values, val)
			}
		}

		if len(params) > 1 {
			interval, _ = strconv.Atoi(params[1])
			return values, &interval
		} else {
			return values, nil
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

// Returns which type of recurrence is: daily/weekly/monthly/yearly
func checkRecurrenceType(weekDays, dates, daily, months []int, weekInterval, datesInterval, dailyInterval, monthInterval *int) string {
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
	// regex := "00 16 * */10 *" // Daily
	// regex := "00 16 1,2,6/2 * *" // Weekly
	regex := "00 16 * (29-31) */1" // Monthly
	// regex := "00 16 * 28 2/12" // Yearly

	regexParams := strings.Split(regex, " ")

	weekDays, weekInterval := getParams(regexParams[2])
	dates, datesInterval := getParams(regexParams[3])
	months, monthInterval := getParams(regexParams[4])

	hours, _ := strconv.Atoi(regexParams[1])
	minutes, _ := strconv.Atoi(regexParams[0])

	endDate, _ := time.Parse(time.RFC3339, "2023-04-23T10:00:00Z")
	startDate, _ := time.Parse(time.RFC3339, "2022-01-25T15:00:00Z")
	currentDate, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	if currentDate.Sub(startDate) > 0 {
		// startDate = currentDate
	}

	recurrenceType := checkRecurrenceType(weekDays, dates, dates, months, weekInterval, datesInterval, datesInterval, monthInterval)
	fmt.Println(recurrenceType)

	switch recurrenceType {
	case "daily":
		nextDues := dailyFunc.GetDailyDues(hours, minutes, *datesInterval, startDate, endDate)
		fmt.Println(nextDues)
	case "weekly":
		// Return only the immediate first occurrence of the each given days of the week
		nextDues := Weekly.GetFirstDueDate(hours, minutes, *weekInterval, startDate, endDate, weekDays)
		// Below function takes the first occurrence of each given week days and
		// returns remaining ones at max 10 or till endDate (whichever is smaller)
		finalDues := Weekly.GetNextDues(*weekInterval, nextDues, endDate)
		fmt.Println(finalDues)
	case "monthly":
		// Return only the immediate first occurrence of the each given dates (monthly interval)
		nextDues := Monthly.GetFirstDateOccurrence(hours, minutes, *monthInterval, startDate, endDate, dates)
		// Below function takes the first occurrence of each given dates and
		// returns remaining ones at max 10 or	till endDate (whichever is smaller)
		finalDues := Monthly.GetNextDues(*monthInterval, endDate, nextDues)
		sort.Sort(finalDues)

		fmt.Println(finalDues)
	case "yearly":
		// Return only the immediate first occurrence of the each given dates (yearly interval)
		nextDues := Yearly.GetFirstDateOccurrence(hours, minutes, *monthInterval, startDate, endDate, dates, months)
		// Below function takes the first occurrence of each given dates and
		// returns remaining ones at max 10 or	till endDate (whichever is smaller)
		finalDues := Yearly.GetNextDues(*monthInterval, endDate, nextDues)
		sort.Sort(finalDues)

		fmt.Println(finalDues)
	}
}
