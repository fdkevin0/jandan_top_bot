package types

import (
	"strconv"
	"time"
)

type Time struct {
	time.Time
}

func (t *Time) UnmarshalJSON(bytes []byte) error {
	ot, err := time.Parse("2006-01-02 15:04:05", string(bytes))
	if err != nil {
		return err
	}
	*t = Time{Time: ot}
	return nil
}

type Timestamp struct {
	time.Time
}

func (t *Timestamp) UnmarshalJSON(bytes []byte) error {
	ts, err := strconv.ParseInt(string(bytes), 10, 64)
	if err != nil {
		return err
	}
	ot := time.Unix(ts, 0)
	if err != nil {
		return err
	}
	*t = Timestamp{Time: ot}
	return nil
}
