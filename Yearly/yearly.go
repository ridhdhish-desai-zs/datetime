package Yearly

import (
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

	count := 0

	for {
		skip := 0
		for i := 0; i < len(firstDateOccurrence); i++ {
			t := firstDateOccurrence[i]
			if count == 10-len(firstDateOccurrence) {
				return nextDues
			}

			nextDate := t.AddDate(0, interval, 0)

			if endDate.Sub(nextDate) < 0 {
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

// Returns first occurrence
// Here startDate contains currentDate or given task startDate (whichever is greater)
func GetFirstDateOccurrence(hrs, mins, interval int, startDate, endDate time.Time, dates, months []int) timeSlice {
	var firstDateOccurrence timeSlice

	for _, v := range dates {
		for _, m := range months {
			nextDate := Monthly.GetMonthDate(hrs, mins, v, m, interval, startDate, endDate)

			if nextDate != nil {
				firstDateOccurrence = append(firstDateOccurrence, *nextDate)
			}
		}
	}

	sort.Sort(firstDateOccurrence)

	return firstDateOccurrence
}
