package utils

import "time"

// TZStringToNative converts a timezone string to *time.Location.
func TZStringToNative(tz string) *time.Location {
	loc, err := time.LoadLocation(tz)
	if err != nil {
		return time.UTC
	}
	return loc
}

// UnixTimestampToStrInTimeZone converts a Unix timestamp to a formatted string in the given timezone.
func UnixTimestampToStrInTimeZone(ts int64, tz *time.Location, format string) string {
	if format == "" {
		format = time.RFC3339
	}
	return time.Unix(ts, 0).In(tz).Format(format)
}
