package main

import (
	"regexp"
)

func CalculatePoints(receipt *Receipt) int {
	return PointsFromAlphanumerics(receipt) +
		pointsFromRoundDollarAmount(receipt) +
		pointsFromQuarters(receipt) +
		pointsFromItemPairs(receipt) +
		pointsFromItemDescriptionLength(receipt) +
		pointsFromPurchaseDayBeingOdd(receipt) +
		pointsFromPurchaseTimeBetween2And4(receipt)
}

// One point for every alphanumeric character in the retailer name.
func PointsFromAlphanumerics(receipt *Receipt) int {
	regex := regexp.MustCompile("[^a-zA-Z0-9]")
	sanitized := regex.ReplaceAllString(receipt.Retailer, "")
	return len(sanitized)
}

// 50 points if the total is a round dollar amount with no cents.
func pointsFromRoundDollarAmount(receipt *Receipt) int {
	return 0
}

// 25 points if the total is a multiple of 0.25.
func pointsFromQuarters(receipt *Receipt) int {
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
