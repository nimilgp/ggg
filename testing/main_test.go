package main

import (
	"testing"
)

func isEven(n int) bool {
	return n%2 == 0
}

func TestFunc(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want bool
	}{
		{"Even number", 4, true},
		{"Odd number", 3, false},
		{"Zero", 0, true},
		{"Negative even", -2, true},
		{"Negative odd", -3, false},
		{"Large even", 1000000, true},
		{"Large odd", 1000001, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEven(tt.num); got != tt.want {
				t.Errorf("isEven(%d) = %v; want %v", tt.num, got, tt.want)
			}
		})
	}
}
