package timehelper

import (
	"time"
)

func FormatYMDHIS(timing *time.Time) string {
	return timing.Format("2006-01-02 15:04:05")
}
