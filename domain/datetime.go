package domain

import "time"

const (
	DATE_TIME_FORMAT_YEAR       = "2006"
	DATE_TIME_FORMAT_YEAR_MONTH = "2006-01"
	DATE_TIME_FORMAT_DATE       = "2006-01-02"
	DATE_TIME_FORMAT_ISO_DATE   = "2006-01-02 15:04:05"
	TIME_TRUNCATION             = 24 * time.Hour
)

func NewDateOnly() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
}

func NewUTCDateOnly() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
}

func ParseOnlyTheDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
}
