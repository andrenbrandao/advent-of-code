package main

import (
	"reflect"
	"testing"
)

func TestDay1(t *testing.T) {
	t.Run("extracts calibration values", func(t *testing.T) {
		calibrationDocument := []string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}
		trebuchet := Trebuchet{calibrationDocument}

		got := trebuchet.CalibrationValues()
		want := []int{12, 38, 15, 77}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
