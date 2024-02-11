package meetup

import "time"

// WeekSchedule defines the type for schedule descriptors.
type WeekSchedule int

const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Fifth
	Last
	Teenth
)

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	// Get the first day of the month
	startDate := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	var targetDate time.Time

	switch wSched {
	case First, Second, Third, Fourth, Fifth:
		for day := 1; day <= 31; day++ {
			currentDate := startDate.AddDate(0, 0, day-1)
			if currentDate.Month() != month {
				break // Moved to the next month
			}
			if currentDate.Weekday() == wDay {
				if int(wSched) == int((currentDate.Day()-1)/7) {
					return currentDate.Day()
				}
			}
		}
	case Last:
		endDate := startDate.AddDate(0, 1, -1) // Last day of the month
		for day := endDate.Day(); day >= 1; day-- {
			currentDate := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
			if currentDate.Weekday() == wDay {
				return currentDate.Day()
			}
		}
	case Teenth:
		for day := 13; day <= 19; day++ {
			currentDate := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
			if currentDate.Weekday() == wDay {
				return currentDate.Day()
			}
		}
	}

	return targetDate.Day()
}
