package time

import (
	"testing"
	"time"
)

func Test_keywordToDate(t *testing.T) {
	tests := []struct {
		name string
		word string
		want string
	}{
		{"today", "today", time.Now().Format(shortForm)},
		{"today mixed case", "Today", time.Now().Format(shortForm)},
		{"yesterday", "yesterday", time.Now().AddDate(0, 0, -1).Format(shortForm)},
		{"yesterday mixed case", "Yesterday", time.Now().AddDate(0, 0, -1).Format(shortForm)},
		{"easter", "2018-04-01", "2018-04-01"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KeywordToDate(tt.word); got != tt.want {
				t.Errorf("keywordToDate() = %v, want %v", got, tt.want)
			}
		})
	}
}
