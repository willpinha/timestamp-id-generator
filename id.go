package main

import "time"

const (
	DefaultValueFormat = "20060102150405"      // YYYYMMDDHHMMSS
	PrettyValueFormat  = "2006-01-02 15:04:05" // YYYY-MM-DD HH:MM:SS
)

type Id struct {
	time  time.Time
	value string
}

func (id Id) Value() string {
	if id.value == "" {
		return id.time.Format(DefaultValueFormat)
	}

	return id.value
}

func (id Id) ValueWithFormat(format string) string {
	return id.time.Format(format)
}

func (id Id) Time() time.Time {
	return id.time
}

func NewId() Id {
	return Id{
		time: time.Now().UTC(),
	}
}
