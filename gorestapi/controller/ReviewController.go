package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ocg.com/go/01/lab/gorestapi/model"
	repo "github.com/ocg.com/go/01/lab/gorestapi/repository"
)

func GetAllReviews(c *fiber.Ctx) error {
	return c.JSON(repo.Reviews.GetAllReviews())
}

func GetReviewById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	review, err := repo.Reviews.FindReviewById(int64(id))
	// fmt.Println(review)
	// fmt.Println(err)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	return c.JSON(review)
}

func DeleteReviewById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	err = repo.Reviews.DeleteReviewById(int64(id))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	} else {
		return c.SendString("delete successfully")
	}
}

func CreateReview(c *fiber.Ctx) error {
	review := new(model.Review)

	err := c.BodyParser(&review)
	fmt.Println(err)
	// if error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	if _, err := repo.Products.FindProductById(review.ProductId); err != nil {
		return c.JSON(fiber.Map{
			"message": "That book does not exist",
		})
	}

	reviewId := repo.Reviews.CreateNewReview(review)
	return c.SendString(fmt.Sprintf("New book is created successfully with id = %d", reviewId))

}

func UpdateReview(c *fiber.Ctx) error {
	updatedReview := new(model.Review)

	err := c.BodyParser(&updatedReview)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	err = repo.Reviews.UpdateReview(updatedReview)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	return c.SendString(fmt.Sprintf("Product with id = %d is successfully updated", updatedReview.Id))

}

func UpsertReview(c *fiber.Ctx) error {
	review := new(model.Review)

	err := c.BodyParser(&review)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	id := repo.Reviews.Upsert(review)
	return c.SendString(fmt.Sprintf("Product with id = %d is successfully upserted", id))
}

// func AverageRating(c *fiber.Ctx) error {
// 	listAllReview := repo.Reviews.GetAllReviews()

// 	for _, detailReview := range listAllReview {
// 		// fmt.Println("TSSSS")
// 		id := detailReview.ProductId

// 		book, err := repo.Products.FindProductById(int64(detailReview.ProductId))

// 		if err != nil {
// 			return c.JSON(fiber.Map{
// 				"message": "Not found book for this id",
// 			})
// 		}

// 		result := repo.Reviews.AverageRating()
// 		book.Rating = result[int64(id)]
// 		repo.Products.UpdateProduct(book)
// 		// return c.SendString(fmt.Sprintf("Product with id = %d rating = %.2f",id, result[int64(id)]))
// 	}
// 	return c.SendString("ok")
// 	// return c.JSON(result[int64(id)])
// }
