// Generated ZEROPS sdk

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"time"
)

type DateTime time.Time

func NewDateTime(value time.Time) DateTime {
	return DateTime(value)
}

func NewDateTimeFromUnix(seconds int64, nanos int64) DateTime {
	return DateTime(time.Unix(seconds, nanos))
}

func NewDateTimeNullFromUnix(seconds int64, nanos int64) DateTimeNull {
	return DateTimeNull{
		filled: true,
		value:  NewDateTimeFromUnix(seconds, nanos),
	}
}

func NewDateTimeFromString(value string) (out DateTime, err error) {
	t, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return out, err
	}
	out = NewDateTime(t)

	return
}

func NewDateTimeNullFromString(value string) (out DateTimeNull, err error) {
	if value == "" {
		return DateTimeNull{}, nil
	}
	v, err := NewDateTimeFromString(value)
	if err != nil {
		return out, err
	}
	return DateTimeNull{
		filled: true,
		value:  v,
	}, nil
}

func (parameter DateTime) DateTimeNull() DateTimeNull {
	return NewDateTimeNull(parameter.Native())
}

func (parameter DateTime) Native() time.Time {
	return time.Time(parameter)
}

func (parameter DateTime) Unix() int64 {
	return time.Time(parameter).Unix()
}

func (parameter DateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(parameter))
}

func (parameter *DateTime) UnmarshalJSON(data []byte) error {
	var value *time.Time
	err := json.NewDecoder(bytes.NewReader(data)).Decode(&value)
	if err != nil {
		return errors.New("value is not valid date format, " + err.Error())
	}
	if value != nil {
		*parameter = DateTime(*value)
	}
	return nil
}

// wrapped original functions

func (parameter DateTime) Add(d time.Duration) DateTime {
	return DateTime(time.Time(parameter).Add(d))
}

func (parameter DateTime) After(u DateTime) bool {
	return time.Time(parameter).After(u.Native())
}

func (parameter DateTime) Before(u DateTime) bool {
	return time.Time(parameter).Before(u.Native())
}

func (parameter DateTime) Sub(u DateTime) time.Duration {
	return time.Time(parameter).Sub(u.Native())
}

func (parameter DateTime) Equal(u DateTime) bool {
	return time.Time(parameter).Equal(u.Native())
}

func (parameter DateTime) Truncate(d time.Duration) DateTime {
	return DateTime(time.Time(parameter).Truncate(d))
}

func (parameter DateTime) Year() int {
	return time.Time(parameter).Year()
}

func (parameter DateTime) Month() time.Month {
	return time.Time(parameter).Month()
}

func (parameter DateTime) Day() int {
	return time.Time(parameter).Day()
}

func (parameter DateTime) Hour() int {
	return time.Time(parameter).Hour()
}

func (parameter DateTime) Minute() int {
	return time.Time(parameter).Minute()
}

func (parameter DateTime) Second() int {
	return time.Time(parameter).Second()
}

func (parameter DateTime) Nanosecond() int {
	return time.Time(parameter).Nanosecond()
}

func (parameter DateTime) String() string {
	return time.Time(parameter).String()
}

func (parameter DateTime) Format(layout string) string {
	return time.Time(parameter).Format(layout)
}
