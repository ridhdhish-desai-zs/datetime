package Monthly

import (
	"fmt"
	"reflect"
	"sort"
	"time"
)

type timeSlice []time.Time

func (s timeSlice) Less(i, j int) bool { return s[i].Before(s[j]) }
func (s timeSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s timeSlice) Len() int           { return len(s) }

/*
	this function will check whether current month has the date which is given in regex input
	if not it will get next occurrence based on interval until it finds valid month or
	endDate exceed
*/
func GetDate(interval, day int, monthDays map[int]int, startDate, endDate time.Time) time.Time {
	var nextDate time.Time
	currentDate := startDate

	for i := 1; ; i++ {
		month := (int(startDate.Month()) + interval*i) % 12
		if month == 0 {
			month = 12
		}
		// checking if month has the input date in it or not
		if monthDays[month] >= day {
			nextDate = startDate.AddDate(0, interval*i, 0)

			return nextDate
		} else {
			currentDate = currentDate.AddDate(0, interval*i, 0)
			if endDate.Sub(currentDate) < 0 {
				return time.Time{}
			}
		}
	}
}

// Calculate first occurrence of the date, based on startDate and endDate and interval
func GetMonthDate(hrs, mins, day, month, interval int, startDate, endDate time.Time) *time.Time {
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

	var nextDate time.Time

	diffOfDays := day - int(startDate.Day())

	// if current date is same as date given in input
	if diffOfDays == 0 {
		minutes := startDate.Hour()*60 + startDate.Minute()
		requiredMinutes := hrs*60 + mins

		// if required time is less than startDate time then it will calculate
		// for next occurrence based on interval
		// input time: 16:00, startDate/current date time: 17:00
		// (startDate time is greater than regex/input time) -> then it will set dueDate of next interval
		if minutes > requiredMinutes {
			nextDate = GetDate(interval, day, monthDays, startDate, endDate)
			if reflect.DeepEqual(nextDate, time.Time{}) {
				return nil
			}
		} else {
			nextDate = startDate
		}
		// if current date is of future
	} else if diffOfDays > 0 {
		if monthDays[int(startDate.Month())] >= day {
			nextDate = startDate.AddDate(0, 0, diffOfDays)
		} else {
			nextDate = GetDate(interval, day, monthDays, startDate, endDate)
			nextDate = nextDate.AddDate(0, 0, diffOfDays)

			if reflect.DeepEqual(nextDate, time.Time{}) {
				return nil
			}
		}
		// if current date is already passed the date given in regex input
	} else {
		nextDate = GetDate(interval, day, monthDays, startDate, endDate)
		nextDate = nextDate.AddDate(0, 0, diffOfDays)

		if reflect.DeepEqual(nextDate, time.Time{}) {
			return nil
		}
	}

	// calculating time after date has been calculated
	nextDate = nextDate.Add(time.Second * time.Duration(0-nextDate.Second()))
	nextDate = nextDate.Add(time.Minute * time.Duration(mins-nextDate.Minute()))
	nextDate = nextDate.Add(time.Hour * time.Duration(hrs-nextDate.Hour()))

	// checking if it does not exceed endDate
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

	// Gives max 10 occurrences or till endDate
	for {
		skip := 0
		for i := 0; i < len(firstDateOccurrence); i++ {
			t := firstDateOccurrence[i]
			if count == 10-len(firstDateOccurrence) {
				return nextDues
			}

			day := t.Day()
			nextDate := GetDate(interval, day, monthDays, t, endDate)

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
