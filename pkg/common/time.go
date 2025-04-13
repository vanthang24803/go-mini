package common

import "time"

func CompareWithNow(expiresAt time.Time) bool {
	return expiresAt.Before(time.Now())
}
