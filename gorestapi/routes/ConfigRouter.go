package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ocg.com/go/01/lab/gorestapi/controller"
)

func ConfigProductRouter(router *fiber.Router) {
	//Return all books
	(*router).Get("/", controller.GetAllProduct) //Liệt kê

	(*router).Get("/:id", controller.GetProductById) //Xem chi tiết một bản ghi

	(*router).Delete("/:id", controller.DeleteProductById) //Xoá một bản ghi

	(*router).Post("", controller.CreateProduct) //INSERT: Tạo một bản ghi

	(*router).Patch("", controller.UpdateProduct) //UPDATE: Cập nhật một bản ghi

	(*router).Put("", controller.UpsertProduct) //UPSERT: Cập nhật một bản ghi nếu tìm thấy còn không tạo mới
}

func ConfigReviewRouter(router *fiber.Router) {
	(*router).Get("/", controller.GetAllReviews)

	(*router).Get("/:id", controller.GetReviewById)

	(*router).Delete("/:id", controller.DeleteReviewById)

	(*router).Post("", controller.CreateReview)

	(*router).Patch("", controller.UpdateReview)

	(*router).Put("", controller.UpsertReview)

	// (*router).Get("/average/", controller.AverageRating)
}

func ConfigCartRouter(router *fiber.Router) {
	(*router).Get("/", controller.GetAllCart)

	(*router).Get("/:id", controller.GetCartById)

	(*router).Delete("/:id", controller.DeleteCartById)

	(*router).Post("", controller.CreateCart)

	(*router).Patch("", controller.UpdateCart)

	(*router).Put("", controller.UpsertCart)

	// (*router).Get("/average/", controller.AverageRating)
}

func ConfigPriceRouter(router *fiber.Router) {
	(*router).Get("/", controller.GetAllPrice)

	(*router).Get("/:id", controller.GetPriceById)

	(*router).Delete("/:id", controller.DeletePriceById)

	(*router).Post("", controller.CreatePrice)

	(*router).Patch("", controller.UpdatePrice)

	(*router).Put("", controller.UpsertPrice)

	// (*router).Get("/average/", controller.AverageRating)
}

func ConfigCategoryRouter(router *fiber.Router) {
	(*router).Get("/", controller.GetAllCategory)

	(*router).Get("/:id", controller.GetCategoryById)

	(*router).Delete("/:id", controller.DeleteCategoryById)

	(*router).Post("", controller.CreateCategory)

	(*router).Patch("", controller.UpdateCategory)

	(*router).Put("", controller.UpsertCategory)

	// (*router).Get("/average/", controller.AverageRating)
}

func ConfigUserRouter(router *fiber.Router) {
	(*router).Get("/", controller.GetAllUser)

	(*router).Get("/:id", controller.GetUserById)

	(*router).Delete("/:id", controller.DeleteUserById)

	(*router).Post("", controller.CreateUser)

	(*router).Patch("", controller.UpdateUser)

	(*router).Put("", controller.UpsertUser)

	// (*router).Get("/average/", controller.AverageRating)
}
