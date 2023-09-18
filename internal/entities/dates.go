package entities

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/imartinezalberte/go-lingq/internal/utils"
)

const (
	DateOnlyLayout                   = time.DateOnly
	DateTimeLayout                   = time.DateTime
	DateHourMinuteLayout             = "15:04"
	DateYearMonthDayHourMinuteLayout = time.DateOnly + utils.Space + DateHourMinuteLayout
)

type ToTimer interface {
	ToTime() time.Time
}

type (
	DateOnly time.Time // Example: 2007-01-02

	DateTime time.Time // Example: 2007-01-02 15:04:05

	DateHourMinute time.Time // Example: 15:04

	DateYearMonthDayHourMinute time.Time // Example: 2007-01-02 15:04
)

// DateOnly
func (d DateOnly) String() string {
	return d.ToTime().Format(DateOnlyLayout)
}

func (d DateOnly) ToTime() time.Time {
	return time.Time(d)
}

func (d *DateOnly) UnmarshalText(b []byte) error {
	date, err := time.Parse(DateOnlyLayout, string(b))
	if err != nil {
		return err
	}

	*d = DateOnly(date)

	return nil
}

func (d DateOnly) MarshalText() ([]byte, error) {
	return []byte(d.String()), nil
}

func (d *DateOnly) UnmarshalJSON(input []byte) error {
	t, err := UnmarshalDateJSON(input, DateOnlyLayout)
	*d = DateOnly(t)
	return err
}

func (d DateOnly) MarshalJSON() ([]byte, error) {
	return MarshalDateJSON(d)
}

// DateOnly
func (d DateTime) String() string {
	return d.ToTime().Format(DateTimeLayout)
}

func (d DateTime) ToTime() time.Time {
	return time.Time(d)
}

func (d *DateTime) UnmarshalText(b []byte) error {
	date, err := time.Parse(DateTimeLayout, string(b))
	if err != nil {
		return err
	}

	*d = DateTime(date)

	return nil
}

func (d DateTime) MarshalText() ([]byte, error) {
	return []byte(d.String()), nil
}

func (d *DateTime) UnmarshalJSON(input []byte) error {
	t, err := UnmarshalDateJSON(input, DateTimeLayout)
	*d = DateTime(t)
	return err
}

func (d DateTime) MarshalJSON() ([]byte, error) {
	return MarshalDateJSON(d)
}

// DateHourMinute
func (d DateHourMinute) String() string {
	return d.ToTime().Format(DateHourMinuteLayout)
}

func (d DateHourMinute) ToTime() time.Time {
	return time.Time(d)
}

func (d *DateHourMinute) UnmarshalText(b []byte) error {
	date, err := time.Parse(DateHourMinuteLayout, string(b))
	if err != nil {
		return err
	}

	*d = DateHourMinute(date)

	return nil
}

func (d DateHourMinute) MarshalText() ([]byte, error) {
	return []byte(d.String()), nil
}

func (d *DateHourMinute) UnmarshalJSON(input []byte) error {
	t, err := UnmarshalDateJSON(input, DateHourMinuteLayout)
	*d = DateHourMinute(t)
	return err
}

func (d DateHourMinute) MarshalJSON() ([]byte, error) {
	return MarshalDateJSON(d)
}

// DateYearMonthDayHourMinute
func (d DateYearMonthDayHourMinute) String() string {
	return d.ToTime().Format(DateYearMonthDayHourMinuteLayout)
}

func (d DateYearMonthDayHourMinute) ToTime() time.Time {
	return time.Time(d)
}

func (d *DateYearMonthDayHourMinute) UnmarshalText(b []byte) error {
	date, err := time.Parse(DateYearMonthDayHourMinuteLayout, string(b))
	if err != nil {
		return err
	}

	*d = DateYearMonthDayHourMinute(date)

	return nil
}

func (d DateYearMonthDayHourMinute) MarshalText() ([]byte, error) {
	return []byte(d.String()), nil
}

func (d *DateYearMonthDayHourMinute) UnmarshalJSON(input []byte) error {
	t, err := UnmarshalDateJSON(input, DateYearMonthDayHourMinuteLayout)
	*d = DateYearMonthDayHourMinute(t)
	return err
}

func (d DateYearMonthDayHourMinute) MarshalJSON() ([]byte, error) {
	return MarshalDateJSON(d)
}

// helpers
func UnmarshalDateJSON(input []byte, layout string) (time.Time, error) {
	sanitized := strings.Trim(string(input), utils.Quote)
	if strings.TrimSpace(sanitized) == "" {
		return time.Time{}, nil
	}

	return time.Parse(layout, sanitized)
}

func MarshalDateJSON[I interface {
	ToTimer
	String() string
}](input I) ([]byte, error) {
	if input.ToTime().IsZero() {
		return json.Marshal(nil)
	}
	return json.Marshal(input.String())
}
