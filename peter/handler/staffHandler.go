package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"peter/staff"
)

func StaffRoutes(app *fiber.App) {
	app.Post("staff/", staffAdd)
}

func staffAdd(ctx *fiber.Ctx) error {
	var s staff.Staff
	err := ctx.BodyParser(&s)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return err
	}
	fmt.Println(s)
	ctx.Status(http.StatusOK)
	return nil
}
