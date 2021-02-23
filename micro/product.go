package micro

import (
	"encoding/json"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/ryanda/product-api-test/service"
)

func Product(c *fiber.Ctx) error {
	result := service.FetchProductData()

	var response interface{}
	err := json.NewDecoder(strings.NewReader(result)).Decode(&response)
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
