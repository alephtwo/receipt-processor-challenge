package main

import (
	"encoding/json"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var receipt1 *Receipt = unmarshalToReceipt(example1)
var receipt2 *Receipt = unmarshalToReceipt(example2)

func TestCalculatePointsExample1(t *testing.T) {
	points := CalculatePoints(receipt1)
	assert.Equal(t, 28, points)
}

func TestCalculatePointsExample2(t *testing.T) {
	points := CalculatePoints(receipt2)
	assert.Equal(t, 109, points)
}

func TestPointsFromAlphanumericsExample1(t *testing.T) {
	points := PointsFromAlphanumerics(receipt1.Retailer)
	assert.Equal(t, 6, points)
}

func TestPointsFromAlphanumericsExample2(t *testing.T) {
	points := PointsFromAlphanumerics(receipt2.Retailer)
	assert.Equal(t, 14, points)
}

func TestPointsFromRoundDollarAmountExample1(t *testing.T) {
	total, _ := strconv.ParseFloat(receipt1.Total, 64)
	points := PointsFromRoundDollarAmount(total)
	assert.Equal(t, 0, points)
}

func TestPointsFromRoundDollarAmountExample2(t *testing.T) {
	total, _ := strconv.ParseFloat(receipt2.Total, 64)
	points := PointsFromRoundDollarAmount(total)
	assert.Equal(t, 50, points)
}

func TestPointsFromQuartersExample1(t *testing.T) {
	total, _ := strconv.ParseFloat(receipt1.Total, 64)
	points := PointsFromQuarters(total)
	assert.Equal(t, 0, points)
}

func TestPointsFromQuartersExample2(t *testing.T) {
	total, _ := strconv.ParseFloat(receipt2.Total, 64)
	points := PointsFromQuarters(total)
	assert.Equal(t, 25, points)
}

func TestPointsFromItemPairsExample1(t *testing.T) {
	points := PointsFromItemPairs(receipt1.Items)
	assert.Equal(t, 10, points)
}

func TestPointsFromItemPairsExample2(t *testing.T) {
	points := PointsFromItemPairs(receipt2.Items)
	assert.Equal(t, 10, points)
}

func TestPointsFromItemPairsThreeItems(t *testing.T) {
	items := make([]Item, 3)
	points := PointsFromItemPairs(items)
	assert.Equal(t, 5, points)
}

func TestPointsFromItemPairsFiveItems(t *testing.T) {
	items := make([]Item, 7)
	points := PointsFromItemPairs(items)
	assert.Equal(t, 15, points)
}

func TestPointsFromItemDescriptionLengthExample1(t *testing.T) {
	points := PointsFromItemDescriptionLength(receipt1.Items)
	assert.Equal(t, 6, points)
}

func TestPointsFromItemDescriptionLengthExample2(t *testing.T) {
	points := PointsFromItemDescriptionLength(receipt2.Items)
	assert.Equal(t, 0, points)
}

func TestPointsFromPurchaseDayBeingOddExample1(t *testing.T) {
	points := PointsFromPurchaseDayBeingOdd(receipt1.PurchaseDate)
	assert.Equal(t, 6, points)
}

func TestPointsFromPurchaseDayBeingOddExample2(t *testing.T) {
	points := PointsFromPurchaseDayBeingOdd(receipt2.PurchaseDate)
	assert.Equal(t, 0, points)
}

func TestPointsFromPurchaseTimeBetween2And4Example1(t *testing.T) {
	points := PointsFromPurchaseTimeBetween2And4(receipt1.PurchaseTime)
	assert.Equal(t, 0, points)
}

func TestPointsFromPurchaseTimeBetween2And4Example2(t *testing.T) {
	points := PointsFromPurchaseTimeBetween2And4(receipt2.PurchaseTime)
	assert.Equal(t, 10, points)
}

func unmarshalToReceipt(input string) *Receipt {
	receipt := new(Receipt)
	// Intentionally ignoring errors, this is a unit test helper.
	// It doesn't need to be quite so robust.
	json.Unmarshal([]byte(input), &receipt)
	return receipt
}

var example1 string = `{
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

var example2 string = `{
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
