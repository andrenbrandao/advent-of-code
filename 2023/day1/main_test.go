package main

import (
	"reflect"
	"testing"
)

func TestDay1(t *testing.T) {
	t.Run("extracts calibration values", func(t *testing.T) {
		calibrationDocument := []string{"1abc2"}
		trebuchet := Trebuchet{calibrationDocument}

		got := trebuchet.CalibrationValues()
		want := []int{12}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
