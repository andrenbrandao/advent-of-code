package trebuchet

import (
	"reflect"
	"testing"
)

func TestCalibrationValues(t *testing.T) {
	t.Run("extracts calibration values", func(t *testing.T) {
		calibrationDocument := []string{
			"1abc2",
			"pqr3stu8vwx",
			"a1b2c3d4e5f",
			"treb7uchet",
		}
		trebuchet := Trebuchet{calibrationDocument}

		got := trebuchet.CalibrationValues()
		want := []int{12, 38, 15, 77}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("extracts words as values", func(t *testing.T) {
		calibrationDocument := []string{
			"two1nine",
			"eightwothree",
			"abcone2threexyz",
			"xtwone3four",
			"4nineeightseven2",
			"zoneight234",
			"7pqrstsixteen",
		}
		trebuchet := Trebuchet{calibrationDocument}

		got := trebuchet.CalibrationValues()
		want := []int{29, 83, 13, 24, 42, 14, 76}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestSum(t *testing.T) {
	t.Run("gets sum of calibration values", func(t *testing.T) {
		calibrationDocument := []string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}
		trebuchet := Trebuchet{calibrationDocument}

		got := trebuchet.Sum()
		want := 142

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
