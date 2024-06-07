package utils

import "time"

func Sleep(d time.Duration) bool {
	time.Sleep(d)
	return true
}

func SleepSeconds(sec int) bool {
	return Sleep(time.Duration(sec) * time.Second)
}

func SleepMinutes(min int) bool {
	return Sleep(time.Duration(min) * time.Minute)
}

func SleepHours(hours int) bool {
	return Sleep(time.Duration(hours) * time.Hour)
}
