package timec

import "time"

func TimeStamp(date time.Time) int64 {
	if date.IsZero() {
		return 0
	}

	return date.Unix()
}
