package Daily

import (
	"time"
)

type timeSlice []time.Time

func (s timeSlice) Less(i, j int) bool { return s[i].Before(s[j]) }
func (s timeSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s timeSlice) Len() int           { return len(s) }

func getDailyDates(hrs, mins, interval int, ct time.Time) time.Time {
	var diffOfDays int

	minutes := ct.Hour()*60 + ct.Minute()
	requiredMinutes := hrs*60 + mins
	if minutes > requiredMinutes {
		diffOfDays = interval
	} else {
		diffOfDays = 0
	}

	nextDate := ct.AddDate(0, 0, diffOfDays)
	nextDate = nextDate.Add(time.Second * time.Duration(0-nextDate.Second()))
	nextDate = nextDate.Add(time.Minute * time.Duration(mins-nextDate.Minute()))
	nextDate = nextDate.Add(time.Hour * time.Duration(hrs-nextDate.Hour()))

	return nextDate
}

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
