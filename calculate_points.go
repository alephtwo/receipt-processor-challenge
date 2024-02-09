package main

import (
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
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

func CalculatePoints(receipt *Receipt) (int, error) {
	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		return 0, err
	}

	points := 0
	points += PointsFromAlphanumerics(receipt.Retailer)
	points += PointsFromRoundDollarAmount(total)
	points += PointsFromQuarters(total)
	points += PointsFromItemPairs(receipt.Items)

	pointsFromDesc, err := PointsFromItemDescriptionLength(receipt.Items)
	if err != nil {
		return 0, err
	}
	points += pointsFromDesc

	pointsFromPurchaseDay, err := PointsFromPurchaseDayBeingOdd(receipt.PurchaseDate)
	if err != nil {
		return 0, err
	}
	points += pointsFromPurchaseDay

	pointsFromPurchaseTime, err := PointsFromPurchaseTimeBetween2And4(receipt.PurchaseTime)
	if err != nil {
		return 0, err
	}
	points += pointsFromPurchaseTime

	return points, nil
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
func PointsFromItemDescriptionLength(items []Item) (int, error) {
	total := 0
	for _, item := range items {
		length := len(strings.TrimSpace(item.ShortDescription))
		if length%3 != 0 {
			// the number is not a multiple of 3.
			continue
		}

		price, err := strconv.ParseFloat(item.Price, 64)
		if err != nil {
			return 0, err
		}

		total += int(math.Ceil(price * 0.2))
	}
	return total, nil
}

// 6 points if the day in the purchase date is odd.
func PointsFromPurchaseDayBeingOdd(purchaseDate string) (int, error) {
	t, err := time.Parse(time.DateOnly, purchaseDate)
	if err != nil {
		return 0, err
	}

	if t.Day()%2 == 1 {
		return 6, nil
	}
	return 0, nil
}

// 10 points if the time of purchase is between 2 and 4 pm.
func PointsFromPurchaseTimeBetween2And4(purchaseTime string) (int, error) {
	t, err := time.Parse("15:04", purchaseTime)
	if err != nil {
		return 0, err
	}

	// Requirements are not exactly clear, but I will assume
	// AFTER 2 (inclusive of 2:00pm)
	// BEFORE 4 (exclusive of 4:00pm)
	if t.Hour() >= 14 && t.Hour() < 16 {
		return 10, nil
	}
	return 0, nil
}
