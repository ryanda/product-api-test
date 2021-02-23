package micro

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ryanda/product-api-test/obj"
	"github.com/ryanda/product-api-test/service"
)

func Product(c *fiber.Ctx) error {
	result, err := service.FetchProductData()
	if err != nil {
		return c.Status(404).JSON(&fiber.Map{
			"success": false,
			"data":    nil,
			"message": err.Error(),
		})
	}

	response, err := obj.UnmarshalProductsResponse(result)
	if err != nil {
		return c.Status(404).JSON(&fiber.Map{
			"success": false,
			"data":    nil,
			"message": err.Error(),
		})
	}

	return c.JSON(&fiber.Map{
		"success": false,
		"data":    response,
		"message": "",
	})
}
