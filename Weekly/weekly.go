package Weekly

import (
	"sort"
	"time"
)

type timeSlice []time.Time

func (s timeSlice) Less(i, j int) bool { return s[i].Before(s[j]) }
func (s timeSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s timeSlice) Len() int           { return len(s) }

// Calculate the first occurrence of given week day and returns that date
// (only used to find first occurrence)
func getDate(hrs, mins, day, recurrence, interval int, startDate time.Time) time.Time {
	wd := map[string]int{
		"Sunday":    0,
		"Monday":    1,
		"Tuesday":   2,
		"Wednesday": 3,
		"Thursday":  4,
		"Friday":    5,
		"Saturday":  6,
	}

	currentDay := wd[startDate.Weekday().String()]
	diffOfDays := day - currentDay

	if diffOfDays == 0 {
		minutes := startDate.Hour()*60 + startDate.Minute()
		requiredMinutes := hrs*60 + mins
		if minutes > requiredMinutes {
			diffOfDays = (recurrence * interval) + diffOfDays
		}
	}

	if diffOfDays < 0 {
		diffOfDays = (recurrence * interval) + diffOfDays
	}

	nextDate := startDate.AddDate(0, 0, diffOfDays)
	nextDate = nextDate.Add(time.Second * time.Duration(0-nextDate.Second()))
	nextDate = nextDate.Add(time.Minute * time.Duration(mins-nextDate.Minute()))
	nextDate = nextDate.Add(time.Hour * time.Duration(hrs-nextDate.Hour()))

	return nextDate
}

// Return first occurrence of each given week day(s)
func GetFirstDueDate(hrs, mins, interval int, startDate, endDate time.Time, weekDays []int) timeSlice {
	var nextDues timeSlice

	// Gets first occurrence of each given week days
	for _, v := range weekDays {
		nextDate := getDate(hrs, mins, v, 7, interval, startDate)

		if endDate.Sub(nextDate) < 0 {
			break
		}
		nextDues = append(nextDues, nextDate)
	}

	sort.Sort(nextDues)

	return nextDues
}

// Returns remaining due dates occurrences (other than first due dates occurrences)
func GetNextDues(interval int, firstDues timeSlice, endDate time.Time) timeSlice {
	count := 0
	var allDues timeSlice

	allDues = append(allDues, firstDues...)

	for {
		skip := 0
		for i := 0; i < len(firstDues); i++ {
			t := firstDues[i]
			if count == 13-len(firstDues) {
				return nil
			}

			nextDate := t.AddDate(0, 0, 7*interval)

			// Checking for endDate (endDate should be greater than or equal)
			// if endDate is gte then append that date in slice
			if endDate.Sub(nextDate) < 0 {
				skip++
			} else {
				count++
				firstDues[i] = nextDate
				allDues = append(allDues, firstDues[i])
			}
		}

		if skip == len(firstDues) {
			return allDues
		}
	}
}
