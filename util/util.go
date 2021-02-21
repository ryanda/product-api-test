package util

import (
	"fmt"
	"time"
)

func GetCurrentTimeStamp() string {
	currentTime := time.Now()

	return currentTime.Format("2006-01-02 15:04:05")
}

func GetTimeSince(start time.Time) string {
	const (
		Decisecond = 100 * time.Millisecond
		Day        = 24 * time.Hour
	)
	sign := time.Duration(1)

	current := time.Since(start)

	if current < 0 {
		sign = -1
		current = -current
	}

	current += +Decisecond / 2
	d := sign * (current / Day)
	current = current % Day
	h := current / time.Hour
	current = current % time.Hour
	m := current / time.Minute
	current = current % time.Minute
	s := current / time.Second
	current = current % time.Second
	f := current / Decisecond

	return fmt.Sprintf("%dd %dh %dm %d.%ds", d, h, m, s, f)
}
