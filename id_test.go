package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestId(t *testing.T) {
	year := 2025
	month := time.May
	day := 30
	hour := 05
	minute := 45
	second := 0
	nsecond := 0

	id := Id{
		time: time.Date(year, month, day, hour, minute, second, nsecond, time.UTC),
	}

	require.Equal(t, "20250530054500", id.Value())
}
