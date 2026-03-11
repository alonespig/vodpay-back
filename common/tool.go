package common

import "time"

func GetTodayTimeRange() (time.Time, time.Time) {
	now := time.Now()
	loc := now.Location()

	start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
	end := start.Add(24 * time.Hour)

	return start, end
}
