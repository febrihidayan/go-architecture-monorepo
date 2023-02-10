package utils

import "time"

func TimeUTC() time.Time {
	return time.Now().UTC()
}
