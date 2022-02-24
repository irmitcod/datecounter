package main

import (
	"datecounter/utils"
	"testing"
)

func TestDate(t *testing.T) {

	tests := []struct {
		from     string
		to       string
		expected int
	}{

		{"2/6/1984", "22/6/1984", 19},
		{"2/6/1983", "22/6/1983", 19},
		{"4/7/1984", "25/12/1984", 173},
		{"3/1/1989", "3/8/1983", 2036},
	}
	utils.UserLeapYear = false
	for _, test := range tests {
		day := utils.IsValidDate(test.from, test.to)
		if got, want := day, test.expected; got != want {
			t.Errorf("Format(%#v,date) = %#v; want %#v", test.from, got, want)
		}
	}
}
