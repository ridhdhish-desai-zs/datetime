package Daily

import (
	"time"
)

type timeSlice []time.Time

func (s timeSlice) Less(i, j int) bool { return s[i].Before(s[j]) }
func (s timeSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s timeSlice) Len() int           { return len(s) }

// Returns first occurrence
// Here startDate contains currentDate or given task startDate whichever is greater
func getDailyDates(hrs, mins, interval int, startDate time.Time) time.Time {
	var diffOfDays int

	// if required time is less than startDate time then it will calculate
	// for next occurrence based on interval
	// input time: 16:00, startDate/current date time: 17:00
	// (startDate time is greater than regex/input time) -> then it will set dueDate of next interval
	minutes := startDate.Hour()*60 + startDate.Minute()
	requiredMinutes := hrs*60 + mins
	if minutes > requiredMinutes {
		diffOfDays = interval
	} else {
		diffOfDays = 0
	}

	nextDate := startDate.AddDate(0, 0, diffOfDays)
	nextDate = nextDate.Add(time.Second * time.Duration(0-nextDate.Second()))
	nextDate = nextDate.Add(time.Minute * time.Duration(mins-nextDate.Minute()))
	nextDate = nextDate.Add(time.Hour * time.Duration(hrs-nextDate.Hour()))

	return nextDate
}

// Calculate next occurrences (other than first occurrences) and returns all possible occurrences
func GetDailyDues(hrs, mins, interval int, startDate, endDate time.Time) []time.Time {
	nextDate := getDailyDates(hrs, mins, interval, startDate)

	var nextDues timeSlice

	if endDate.Sub(nextDate) < 0 {
		return nextDues
	}

	nextDues = append(nextDues, nextDate)

	for i := 0; i < 10; i++ {
		nextDate = nextDate.AddDate(0, 0, interval)

		if endDate.Sub(nextDate) < 0 {
			return nextDues
		}

		nextDues = append(nextDues, nextDate)
	}

	return nextDues
}
