package Monthly

import (
	"fmt"
	"sort"
	"time"
)

type timeSlice []time.Time

func (s timeSlice) Less(i, j int) bool { return s[i].Before(s[j]) }
func (s timeSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s timeSlice) Len() int           { return len(s) }

// Calculate first occurrence of the date, based on startDate and endDate and interval
func GetMonthDate(hrs, mins, day, month, interval int, startDate, endDate time.Time) *time.Time {
	var nextDate time.Time
	nextDate = startDate

	var diffOfMonths int

	diffOfDays := day - int(startDate.Day())
	if month != 0 {
		diffOfMonths = month - int(startDate.Month())
	}

	// if required time is less than startDate time then it will calculate
	// for next occurrence based on interval
	// input time: 16:00, startDate/current date time: 17:00
	// (startDate time is greater than regex/input time) -> then it will set dueDate of next interval
	if diffOfMonths == 0 {
		if diffOfDays == 0 {
			minutes := startDate.Hour()*60 + startDate.Minute()
			requiredMinutes := hrs*60 + mins
			if minutes > requiredMinutes {
				nextDate = startDate.AddDate(0, interval, 0)
			}
		}

		if diffOfDays < 0 {
			nextDate = startDate.AddDate(0, interval, 0)
		}
	} else if diffOfMonths < 0 {
		nextDate = startDate.AddDate(0, interval, 0)
	}

	nextDate = nextDate.AddDate(0, diffOfMonths, diffOfDays)
	nextDate = nextDate.Add(time.Second * time.Duration(0-nextDate.Second()))
	nextDate = nextDate.Add(time.Minute * time.Duration(mins-nextDate.Minute()))
	nextDate = nextDate.Add(time.Hour * time.Duration(hrs-nextDate.Hour()))

	if endDate.Sub(nextDate) < 0 {
		return nil
	}

	return &nextDate
}

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
func GetFirstDateOccurrence(hrs, mins, interval int, startDate, endDate time.Time, dates []int) timeSlice {
	var firstDateOccurrence timeSlice

	for _, date := range dates {
		nextDate := GetMonthDate(hrs, mins, date, 0, interval, startDate, endDate)

		if nextDate != nil {
			firstDateOccurrence = append(firstDateOccurrence, *nextDate)
		}
	}

	sort.Sort(firstDateOccurrence)

	fmt.Println(firstDateOccurrence)

	return firstDateOccurrence
}
