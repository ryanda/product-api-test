package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ryanda/product-api-test/config"
	"github.com/ryanda/product-api-test/micro"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.InitApp()

	setupOnCloseHandler()

	app := fiber.New(fiber.Config{
		ErrorHandler: micro.ErrHandler,
	})

	app.Get("/", micro.Index)
	app.Get("/product", micro.Product)

	app.Use(micro.NotFoundHandler)

	// Listen on port 3000
	log.Fatal(app.Listen(":" + config.ServicePort)) // go run app.go -port=3000
}

func setupOnCloseHandler() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Service stopped!")
		os.Exit(0)
	}()
}
