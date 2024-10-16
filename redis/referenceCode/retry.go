package referenceCode

import "time"

type RetryStrategy interface {
	// 返回下一次重试的间隔，如果不需要重试，第二参数返回 false
	Next() (time.Duration, bool)
}

type FixIntervalRetry struct {
	// 最大重试间隔
	Interval time.Duration
	// 最大次数
	Max int
	cnt int
}

func (f *FixIntervalRetry) Next() (time.Duration, bool) {
	f.cnt++
	return f.Interval, f.cnt <= f.Max
}
