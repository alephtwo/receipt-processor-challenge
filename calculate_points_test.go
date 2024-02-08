package main

import (
	"encoding/json"
	"testing"
)

func TestCalculatePointsExample1(t *testing.T) {
	receipt := unmarshalToReceipt(example_1)
	points := CalculatePoints(receipt)
	if points != 28 {
		t.Fatalf("Expected 28 points, got %d", points)
	}
}

func TestCalculatePointsExample2(t *testing.T) {
	receipt := unmarshalToReceipt(example_2)
	points := CalculatePoints(receipt)
	if points != 100 {
		t.Fatalf("Expected 100 points, got %d", points)
	}
}

func TestPointsFromAlphanumericsExample1(t *testing.T) {
	receipt := unmarshalToReceipt(example_1)
	points := PointsFromAlphanumerics(receipt)
	if points != 6 {
		t.Fatalf("Expected 6 points, got %d", points)
	}
}

func TestPointsFromAlphanumericsExample2(t *testing.T) {
	receipt := unmarshalToReceipt(example_2)
	points := PointsFromAlphanumerics(receipt)
	if points != 14 {
		t.Fatalf("Expected 14 points, got %d", points)
	}
}

func unmarshalToReceipt(input string) *Receipt {
	receipt := new(Receipt)
	// Intentionally ignoring errors, this is a unit test helper.
	// It doesn't need to be quite so robust.
	json.Unmarshal([]byte(input), &receipt)
	return receipt
}

var example_1 string = `{
  "retailer": "Target",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "items": [
    {
      "shortDescription": "Mountain Dew 12PK",
      "price": "6.49"
    },{
      "shortDescription": "Emils Cheese Pizza",
      "price": "12.25"
    },{
      "shortDescription": "Knorr Creamy Chicken",
      "price": "1.26"
    },{
      "shortDescription": "Doritos Nacho Cheese",
      "price": "3.35"
    },{
      "shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
      "price": "12.00"
    }
  ],
  "total": "35.35"
}`

var example_2 string = `{
  "retailer": "M&M Corner Market",
  "purchaseDate": "2022-03-20",
  "purchaseTime": "14:33",
  "items": [
    {
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    }
  ],
  "total": "9.00"
}`
