package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	tz, _ := time.LoadLocation("Asia/Tokyo")
	future := time.Date(2015, time.October, 21, 7, 28, 0, 0, tz)

	fmt.Println(now.String())
	fmt.Println(future.Format(time.RFC3339Nano))
	fmt.Println(NextMonth(time.Date(2021, 5, 31, 0, 0, 0, 0, time.UTC)).String())
}

func NextMonth(t time.Time) time.Time {
	year1, month1, day := t.Date()
	first := time.Date(year1, month1, 1, 0, 0, 0, 0, time.UTC)
	year2, month2, _ := first.AddDate(0, 1, 0).Date()
	nextMonthTime := time.Date(year2, month2, day, 0, 0, 0, 0, time.UTC)
	if month2 != nextMonthTime.Month() {
		return first.AddDate(0, 2, -1) // 翌月末
	}
	return nextMonthTime
}
