package main

import (
	"log"
	"math"
	"regexp"
	"strconv"
)

func CalculatePoints(receipt *Receipt) int {
	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		log.Println("Total was not a valid dollar value.")
		return 0
	}

	return PointsFromAlphanumerics(receipt.Retailer) +
		PointsFromRoundDollarAmount(total) +
		pointsFromQuarters(total) +
		pointsFromItemPairs(receipt) +
		pointsFromItemDescriptionLength(receipt) +
		pointsFromPurchaseDayBeingOdd(receipt) +
		pointsFromPurchaseTimeBetween2And4(receipt)
}

// One point for every alphanumeric character in the retailer name.
func PointsFromAlphanumerics(retailer string) int {
	regex := regexp.MustCompile("[^a-zA-Z0-9]")
	sanitized := regex.ReplaceAllString(retailer, "")
	return len(sanitized)
}

// 50 points if the total is a round dollar amount with no cents.
func PointsFromRoundDollarAmount(total float64) int {
	// Could also check this by just checking if the string ends with ".00".
	// That's probably faster, but this is more robust at handling unexpected data.
	//
	// If the total is equal to the value of itself less any decimal positions...
	if total == math.Trunc(total) {
		return 50
	}
	return 0
}

// 25 points if the total is a multiple of 0.25.
func pointsFromQuarters(total float64) int {
	return 0
}

// 5 points for every two items on the receipt.
func pointsFromItemPairs(receipt *Receipt) int {
	return 0
}

// If the trimmed length of the item description is a multiple of 3,
// multiply the price by 0.2 and round up to the nearest integer.
// The result is the number of points earned.
func pointsFromItemDescriptionLength(receipt *Receipt) int {
	return 0
}

// 6 points if the day in the purchase date is odd.
func pointsFromPurchaseDayBeingOdd(receipt *Receipt) int {
	return 0
}

func pointsFromPurchaseTimeBetween2And4(receipt *Receipt) int {
	return 0
}
