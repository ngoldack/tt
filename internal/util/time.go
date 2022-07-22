package util

import "time"

func FormatTime(time time.Time) string {
	return time.Format("15:04:05")
}
