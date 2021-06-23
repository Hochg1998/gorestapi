package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ocg.com/go/01/lab/gorestapi/model"
	repo "github.com/ocg.com/go/01/lab/gorestapi/repository"
)

func GetAllPrice(c *fiber.Ctx) error {
	return c.JSON(repo.Prices.GetAllPrices())
}

func GetPriceById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	price, err := repo.Prices.FindPriceById(int64(id))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	return c.JSON(price)
}

func DeletePriceById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	err = repo.Prices.DeletePriceById(int64(id))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	} else {
		return c.SendString("delete price successfully")
	}
}

func CreatePrice(c *fiber.Ctx) error {
	price := new(model.Price)

	err := c.BodyParser(&price)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	priceId := repo.Prices.CreateNewPrice(price)
	return c.SendString(fmt.Sprintf("New price is created successfully with id = %d", priceId))

}

func UpdatePrice(c *fiber.Ctx) error {
	updatedPrice := new(model.Price)

	err := c.BodyParser(&updatedPrice)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	err = repo.Prices.UpdatePrice(updatedPrice)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	return c.SendString(fmt.Sprintf("Price with id = %d is successfully updated", updatedPrice.Id))

}

func UpsertPrice(c *fiber.Ctx) error {
	price := new(model.Price)

	err := c.BodyParser(&price)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	id := repo.Prices.Upsert(price)
	return c.SendString(fmt.Sprintf("Price with id = %d is successfully upserted", id))
}
