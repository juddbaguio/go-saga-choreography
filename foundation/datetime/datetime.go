package datetime

import (
	"time"
)

const (
	Timezone   = "Asia/Bangkok"
	DateTime   = "2006-01-02 15:04:05"
	Date       = "2006-01-02"
	Time       = "15.04"
	Time2      = "15:04"
	Time3      = "15:04:05" // used for testing purposes
	UTC_OFFSET = "2006-01-02T15:04:05+0700"
)

func DateToIso8601(dateTime string) string {
	return LocalDate(dateTime, DateTime).Format(time.RFC3339)
}

func Iso8601ToDate(dateTime string) string {
	return LocalDate(dateTime, time.RFC3339).Format(Date)
}

func LocalDate(dateTime string, format string) time.Time {
	localLoc, _ := time.LoadLocation(Timezone)
	date, _ := time.Parse(format, dateTime)
	return date.UTC().In(localLoc)
}
func UTCToLocalDate(dateTime string, format string) time.Time {
	localLoc, _ := time.LoadLocation(Timezone)
	date, _ := time.Parse(format, dateTime)
	return date.In(localLoc)
}

func DateToUTC(date string) time.Time {
	loc, _ := time.LoadLocation(Timezone)
	localDate, _ := time.ParseInLocation(Date, date, loc)
	return localDate.UTC()
}

func DateTimeToUTC(dateTime string) time.Time {
	loc, _ := time.LoadLocation(Timezone)
	localDate, _ := time.ParseInLocation(DateTime, dateTime, loc)
	return localDate.UTC()
}

func UTCDateToFilterDate(date string) (startDate time.Time, endDate time.Time) {
	tmpDate, _ := time.Parse(Date, date)
	tmpDate = tmpDate.Add(24 * time.Hour)
	loc, _ := time.LoadLocation(Timezone)
	localDate, _ := time.ParseInLocation(Date, tmpDate.Format(Date), loc)
	startDate = localDate.UTC().Add(-24 * time.Hour)
	endDate = localDate.UTC().Add(-time.Duration(localDate.Hour())).Add(-1 * time.Second)
	return startDate, endDate
}
