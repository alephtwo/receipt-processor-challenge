package main

import (
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// The given OpenAPI spec defines number fields as strings... so I'm trying
// to assume that's what we want both in and out. It makes some things harder.
//
// A lot of the logic here can be done with raw strings.
// For example, checking if the total ends with ".00".
//
// That's probably faster, but this is more robust at handling unexpected data.
// Plus, I feel like it is better to work in the intended formats for
// each given data point. Clearly we are doing math on these values,
// and math should (ideally) be done with numbers!

func CalculatePoints(receipt *Receipt) int {
	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		log.Println("Total was not a valid dollar value.")
		return 0
	}

	return PointsFromAlphanumerics(receipt.Retailer) +
		PointsFromRoundDollarAmount(total) +
		PointsFromQuarters(total) +
		PointsFromItemPairs(receipt.Items) +
		PointsFromItemDescriptionLength(receipt.Items) +
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
	// If the total is equal to the value of itself less any decimal positions...
	if total == math.Trunc(total) {
		return 50
	}
	return 0
}

// 25 points if the total is a multiple of 0.25.
func PointsFromQuarters(total float64) int {
	if math.Mod(total, 0.25) == 0 {
		return 25
	}
	return 0
}

// 5 points for every two items on the receipt.
func PointsFromItemPairs(items []Item) int {
	return (len(items) / 2) * 5
}

// If the trimmed length of the item description is a multiple of 3,
// multiply the price by 0.2 and round up to the nearest integer.
// The result is the number of points earned.
func PointsFromItemDescriptionLength(items []Item) int {
	total := 0
	for _, item := range items {
		length := len(strings.TrimSpace(item.ShortDescription))
		if length%3 != 0 {
			// the number is not a multiple of 3.
			continue
		}

		price, err := strconv.ParseFloat(item.Price, 64)
		if err != nil {
			log.Println("Price was not a valid dollar value.")
			continue
		}

		total += int(math.Ceil(price * 0.2))
	}
	return total
}

// 6 points if the day in the purchase date is odd.
func pointsFromPurchaseDayBeingOdd(receipt *Receipt) int {
	return 0
}

func pointsFromPurchaseTimeBetween2And4(receipt *Receipt) int {
	return 0
}
