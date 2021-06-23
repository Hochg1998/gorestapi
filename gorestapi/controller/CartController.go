package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ocg.com/go/01/lab/gorestapi/model"
	repo "github.com/ocg.com/go/01/lab/gorestapi/repository"
)

func GetAllCart(c *fiber.Ctx) error {
	return c.JSON(repo.Carts.GetAllCarts())
}

func GetCartById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	cart, err := repo.Carts.FindCartById(int64(id))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	return c.JSON(cart)
}

func DeleteCartById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	err = repo.Carts.DeleteCartById(int64(id))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	} else {
		return c.SendString("delete cart successfully")
	}
}

func CreateCart(c *fiber.Ctx) error {
	cart := new(model.Cart)

	err := c.BodyParser(&cart)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	cartId := repo.Carts.CreateNewCart(cart)
	return c.SendString(fmt.Sprintf("New cart is created successfully with id = %d", cartId))

}

func UpdateCart(c *fiber.Ctx) error {
	updatedCart := new(model.Cart)

	err := c.BodyParser(&updatedCart)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	err = repo.Carts.UpdateCart(updatedCart)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	return c.SendString(fmt.Sprintf("Cart with id = %d is successfully updated", updatedCart.Id))

}

func UpsertCart(c *fiber.Ctx) error {
	cart := new(model.Cart)

	err := c.BodyParser(&cart)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	id := repo.Carts.Upsert(cart)
	return c.SendString(fmt.Sprintf("Cart with id = %d is successfully upserted", id))
}
