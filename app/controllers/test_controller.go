package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Test(c *fiber.Ctx) error {
	var err error
	err = nil
	fmt.Println("test")
	return err
}
