package Yearly

import (
	"fmt"
	"reflect"
	"sort"
	"time"

	"github.com/ridhdhish-desai-zs/datetime/Monthly"
)

type timeSlice []time.Time

func (s timeSlice) Less(i, j int) bool { return s[i].Before(s[j]) }
func (s timeSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s timeSlice) Len() int           { return len(s) }

// Calculate next occurrences (other than first occurrences) and returns
// all possible occurrences (at max 10)
func GetNextDues(interval int, endDate time.Time, firstDateOccurrence timeSlice) timeSlice {
	var nextDues timeSlice
	nextDues = append(nextDues, firstDateOccurrence...)

	monthDays := map[int]int{
		1:  31,
		2:  28,
		3:  31,
		4:  30,
		5:  31,
		6:  30,
		7:  31,
		8:  31,
		9:  30,
		10: 31,
		11: 30,
		12: 31,
	}

	count := 0

	for {
		skip := 0
		for i := 0; i < len(firstDateOccurrence); i++ {
			t := firstDateOccurrence[i]
			if count == 10-len(firstDateOccurrence) {
				return nextDues
			}

			day := t.Day()

			nextDate := Monthly.GetDate(interval, day, monthDays, t, endDate)

			if reflect.DeepEqual(nextDate, time.Time{}) || endDate.Sub(nextDate) < 0 {
				skip++
			} else {
				count++
				firstDateOccurrence[i] = nextDate
				nextDues = append(nextDues, firstDateOccurrence[i])
			}
		}

		if skip == len(firstDateOccurrence) {
			return nextDues
		}
	}
}

// TODO: create new function for get first occurrences
// 1. Check for month (occurred, not occurred or same month)
func getMonthDate(hrs, mins, date, month, interval int, startDate, endDate time.Time) *time.Time {
	monthDays := map[int]int{
		1:  31,
		2:  28,
		3:  31,
		4:  30,
		5:  31,
		6:  30,
		7:  31,
		8:  31,
		9:  30,
		10: 31,
		11: 30,
		12: 31,
	}

	diffOfMonths := month - int(startDate.Month())
	diffOfDays := date - startDate.Day()

	var nextDate *time.Time

	if diffOfMonths == 0 {
		nextDate = Monthly.GetMonthDate(hrs, mins, date, month, interval, startDate, endDate)
	} else if diffOfMonths > 0 {

	}

	return nil
}

// Returns first occurrence
// Here startDate contains currentDate or given task startDate (whichever is greater)
func GetFirstDateOccurrence(hrs, mins, interval int, startDate, endDate time.Time, dates, months []int) timeSlice {
	var firstDateOccurrence timeSlice

	for _, v := range dates {
		for _, m := range months {

			// FIXME: Don't use month's function
			nextDate := Monthly.GetMonthDate(hrs, mins, v, m, interval, startDate, endDate)

			if nextDate != nil {
				firstDateOccurrence = append(firstDateOccurrence, *nextDate)
			}
		}
	}

	sort.Sort(firstDateOccurrence)

	fmt.Println(firstDateOccurrence)

	return firstDateOccurrence
}
