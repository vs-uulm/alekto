package time

import "time"

const (
	day = 60 * 60 * 24
)

func NowTimestamp () int64 {
	return time.Now().UnixNano() / int64(time.Second)
}

func YesterdayTimestamp () int64 {
	return NowTimestamp() - 1 * day
}