package lang

import "time"

// Time 用 int64 表示一个 unix 时间戳，单位是毫秒，基于 UTC_1970-01-01_00:00:00
type Time int64

// Seconds 以秒为单位表示时间长度
type Seconds int64

// Milliseconds 以毫秒为单位表示时间长度
type Milliseconds int64

////////////////////////////////////////////////////////////////////////////////

// Sub 函数提供时间戳的减法，计算 从 t0 开始，到 t 为止所经历的时间
func (t Time) Sub(t0 Time) time.Duration {
	ms := Milliseconds(t - t0)
	return ms.Duration()
}

// Add 函数提供时间戳的加法，计算从 t 开始，经历 d 时间后，所在的时刻
func (t Time) Add(d time.Duration) Time {
	ms := NewMilliseconds(d)
	return t + Time(ms)
}

func (t Time) String() string {
	tt := t.Time()
	return tt.String()
}

// Time 把时间戳转换为 time.Time 形式
func (t Time) Time() time.Time {
	if t == 0 {
		var zero time.Time
		return zero
	}
	ms := int64(t)
	return time.UnixMilli(ms)
}

// Int 把时间戳转换为 int64 形式
func (t Time) Int() int64 {
	return int64(t)
}

// NewTime 根据 time.Time 创建时间戳
func NewTime(t time.Time) Time {
	if t.IsZero() {
		return 0
	}
	ms := t.UnixMilli()
	return Time(ms)
}

// Now 取当前时间戳
func Now() Time {
	now := time.Now()
	return NewTime(now)
}

////////////////////////////////////////////////////////////////////////////////

// Duration 转换为 Duration 形式
func (value Seconds) Duration() time.Duration {
	const x = time.Second
	return time.Duration(value) * x
}

// NewSeconds 把 Duration 转换为 Seconds
func NewSeconds(d time.Duration) Seconds {
	const x = time.Second
	return Seconds(d / x)
}

////////////////////////////////////////////////////////////////////////////////

// Duration 转换为 Duration 形式
func (value Milliseconds) Duration() time.Duration {
	const x = time.Millisecond
	return time.Duration(value) * x
}

// NewMilliseconds 把 Duration 转换为 Milliseconds
func NewMilliseconds(d time.Duration) Milliseconds {
	const x = time.Millisecond
	return Milliseconds(d / x)
}

////////////////////////////////////////////////////////////////////////////////
