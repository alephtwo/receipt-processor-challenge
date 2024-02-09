package main

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestExampleReceipt1(t *testing.T) {
	testReceipt(t, exampleReceiptJson1, 28)
}

func TestExampleReceipt2(t *testing.T) {
	testReceipt(t, exampleReceiptJson2, 109)
}

func testReceipt(t *testing.T, receiptJson string, points int) {
	app := CreateApp()
	id := testProcessReceipt(t, app, receiptJson)
	testPoints(t, app, id, points)
}

func testProcessReceipt(t *testing.T, app *fiber.App, receiptJson string) string {
	// Process the receipt
	req := httptest.NewRequest("POST", "/receipts/process", strings.NewReader(receiptJson))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)
	assert.Equal(t, 200, resp.StatusCode)

	// Get the contents of the request
	defer resp.Body.Close()
	responseJson, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf(err.Error())
	}

	var processReceiptResponse ProcessReceiptResponse
	json.Unmarshal(responseJson, &processReceiptResponse)
	assert.NotNil(t, processReceiptResponse.Id)

	return processReceiptResponse.Id
}

func testPoints(t *testing.T, app *fiber.App, id string, points int) {
	req := httptest.NewRequest("GET", "/receipts/"+id+"/points", nil)
	resp, _ := app.Test(req)
	assert.Equal(t, 200, resp.StatusCode)

	// Get the contents of the request
	defer resp.Body.Close()
	responseJson, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf(err.Error())
	}

	var pointsResponse PointsResponse
	json.Unmarshal(responseJson, &pointsResponse)
	assert.Equal(t, points, pointsResponse.Points)
}

func TestPointsInvalidId(t *testing.T) {
	app := CreateApp()
	req := httptest.NewRequest("GET", "/receipts/1a57debd-e627-4e0d-8898-0d98b9a412f1/points", nil)
	resp, _ := app.Test(req)
	assert.Equal(t, 404, resp.StatusCode)
}

type ProcessReceiptResponse struct {
	Id string
}

type PointsResponse struct {
	Points int
}
