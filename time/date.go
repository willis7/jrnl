package time

import "time"

const shortForm = "2006-01-02"

// Today returns a short form string representation for today
func Today() string {
	return time.Now().Format(shortForm)
}

// Yesterday returns a short form string representation for yesterday
func Yesterday() string {
	return time.Now().AddDate(0, 0, -1).Format(shortForm)
}

// ShortFormString converts a time object to a
// short form string e.g. 2018-01-30
func ShortFormString(time *time.Time) string {
	return time.Format(shortForm)
}
