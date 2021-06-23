package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ocg.com/go/01/lab/gorestapi/model"
	repo "github.com/ocg.com/go/01/lab/gorestapi/repository"
)

func GetAllUser(c *fiber.Ctx) error {
	return c.JSON(repo.Users.GetAllUsers())
}

func GetUserById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	user, err := repo.Users.FindUserById(int64(id))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	return c.JSON(user)
}

func DeleteUserById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	err = repo.Users.DeleteUserById(int64(id))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	} else {
		return c.SendString("delete user successfully")
	}
}

func CreateUser(c *fiber.Ctx) error {
	user := new(model.User)

	err := c.BodyParser(&user)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	userId := repo.Users.CreateNewUser(user)
	return c.SendString(fmt.Sprintf("New user is created successfully with id = %d", userId))

}

func UpdateUser(c *fiber.Ctx) error {
	updatedUser := new(model.User)

	err := c.BodyParser(&updatedUser)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	err = repo.Users.UpdateUser(updatedUser)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	return c.SendString(fmt.Sprintf("User with id = %d is successfully updated", updatedUser.Id))

}

func UpsertUser(c *fiber.Ctx) error {
	user := new(model.User)

	err := c.BodyParser(&user)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	id := repo.Users.Upsert(user)
	return c.SendString(fmt.Sprintf("User with id = %d is successfully upserted", id))
}
