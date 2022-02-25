package Yearly

import (
	"sort"
	"time"
)

type timeSlice []time.Time

func (s timeSlice) Less(i, j int) bool { return s[i].Before(s[j]) }
func (s timeSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s timeSlice) Len() int           { return len(s) }

func getMonthDate(hrs, mins, day, month, interval int, startDate, endDate time.Time) *time.Time {
	var nextDate time.Time
	nextDate = startDate

	diffOfDays := day - int(startDate.Day())
	diffOfMonths := month - int(startDate.Month())

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

	if endDate.Sub(nextDate) < 0 || (month != int(nextDate.Month())) {
		return nil
	}

	return &nextDate
}

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

func GetFirstDateOccurrence(hrs, mins, interval int, startDate, endDate time.Time, dates, months []int) timeSlice {
	var firstDateOccurrence timeSlice

	for _, v := range dates {
		for _, m := range months {
			nextDate := getMonthDate(hrs, mins, v, m, interval, startDate, endDate)

			if nextDate != nil {
				firstDateOccurrence = append(firstDateOccurrence, *nextDate)
			}
		}
	}

	sort.Sort(firstDateOccurrence)

	return firstDateOccurrence
}
