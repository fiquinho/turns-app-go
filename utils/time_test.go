package utils

import (
	"reflect"
	"testing"
	"time"
)

func assertDateTime(t *testing.T, expected TimeRange, actual TimeRange) {
	if reflect.DeepEqual(expected, actual) == false {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestGetWeekByDay(t *testing.T) {

	weekRange := TimeRange{
		StartTime: time.Date(2024, 2, 26, 0, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2024, 3, 3, 23, 59, 59, 0, time.UTC),
	}

	t.Run("test week mid day", func(t *testing.T) {
		day := time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC)
		week := GetWeekByDay(day)

		assertDateTime(t, weekRange, week)
	})

	t.Run("test week monday", func(t *testing.T) {
		day := time.Date(2024, 2, 26, 0, 0, 0, 0, time.UTC)
		week := GetWeekByDay(day)

		assertDateTime(t, weekRange, week)
	})

	t.Run("test week sunday", func(t *testing.T) {
		day := time.Date(2024, 3, 3, 23, 59, 59, 0, time.UTC)
		week := GetWeekByDay(day)

		assertDateTime(t, weekRange, week)
	})

}
