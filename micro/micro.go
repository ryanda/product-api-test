package micro

import (
	"fmt"

	"github.com/ryanda/product-api-test/config"
	"github.com/ryanda/product-api-test/util"

	"github.com/gofiber/fiber/v2"
)

func ErrHandler(c *fiber.Ctx, err error) error {
	// Statuscode defaults to 500
	code := fiber.StatusInternalServerError

	// Retreive the custom statuscode if it's an fiber.*Error
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	// Return HTTP response
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

	return c.Status(code).JSON(&fiber.Map{
		"error": err.Error(),
	})
}

func NotFoundHandler(c *fiber.Ctx) error {
	return c.Status(404).JSON(&fiber.Map{
		"success": false,
		"data":    nil,
		"message": "Not found",
	})
}

func Index(c *fiber.Ctx) error {

	start := config.GetTime()
	duration := util.GetTimeSince(start)

	response := fmt.Sprintf("engine start at: %s (%s)", start.Format("2006-01-02 15:04:05"), duration)

	return c.JSON(&fiber.Map{
		"success": true,
		"data":    response,
	})
}
