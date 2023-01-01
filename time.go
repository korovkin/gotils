package gotils

import (
	"strconv"
	"time"
)

type Time time.Time

func Now() Time {
	return Time(time.Now())
}

func FromUnix(sec int64) Time {
	return Time(time.Unix(sec, 0))
}

func FromUnixMilli(sec int64) Time {
	return Time(time.UnixMilli(sec))
}

func (t Time) Time() time.Time {
	return time.Time(t)
}

func (t Time) MarshalJSON() ([]byte, error) {
	unixTime := time.Time(t).Unix()
	// do not go below zero:
	if unixTime <= 0 {
		unixTime = 0
	}
	return []byte(strconv.FormatInt(unixTime, 10)), nil
}

func (t *Time) UnmarshalJSON(s []byte) (err error) {
	r := string(s)
	q, err := strconv.ParseInt(r, 10, 64)
	CheckNotFatal(err)
	if err != nil {
		return err
	}

	if q > 0 {
		*(*time.Time)(t) = time.Unix(q, 0)
	} else {
		*(*time.Time)(t) = time.Time{}
	}

	return nil
}

func (t Time) Unix() int64 {
	return time.Time(t).Unix()
}

func (t Time) UTC() time.Time {
	return time.Time(t).UTC()
}

func (t Time) Local() time.Time {
	return time.Time(t).Local()
}

func (t Time) String() string {
	return t.UTC().String()
}
