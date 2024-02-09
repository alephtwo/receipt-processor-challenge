package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var receipts = make(map[uuid.UUID]int)

func main() {
	app := CreateApp()
	app.Listen(":8080")
}

func CreateApp() *fiber.App {
	app := fiber.New()
	app.Post("/receipts/process", processReceipt)
	app.Get("/receipts/:id/points", getPoints)
	return app
}

func processReceipt(c *fiber.Ctx) error {
	c.Accepts("application/json")

	receipt := new(Receipt)
	if err := c.BodyParser(receipt); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	points, err := CalculatePoints(receipt)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	id := uuid.New()
	receipts[id] = points

	return c.JSON(fiber.Map{"id": id})
}

func getPoints(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	points, ok := receipts[id]
	if !ok {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.JSON(fiber.Map{"points": points})
}

type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}
