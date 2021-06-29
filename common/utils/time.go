package utils

import (
	"gym-app/common/logger"
	"time"
)

var log = logger.NewLogger()

func ParseDate(date string) time.Time {
	if date == "" {
		return time.Now()
	}
	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		log.WithField("err", err).
			Errorf("Couldn't parse date %s", date)
	}
	return parsedDate
}

func StartOfTheDay(date time.Time) time.Time {
	year, month, day := date.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, date.Location())
}

func EndOfTheDay(date time.Time) time.Time {
	year, month, day := date.Date()
	return time.Date(year, month, day, 23, 59, 59, 0, date.Location())
}
