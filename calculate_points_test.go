package main

import (
	"testing"
)

func TestCalculatePoints(t *testing.T) {
	receipt := new(Receipt)

	points := CalculatePoints(receipt)
	if points == 0 {
		t.Fatalf("Points are 0")
	}
}
