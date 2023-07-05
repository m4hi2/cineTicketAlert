package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/m4hi2/capsule71/pkg/dbconn"
)

func New() (*fiber.App, error) {
	if err := dbconn.DoPersistConnect(); err != nil {
		return nil, err
	}

	app := fiber.New()
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("gg")
	})

	return app, nil
}
