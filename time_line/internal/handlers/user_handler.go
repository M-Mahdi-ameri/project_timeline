package handlers

import (
	"context"
	"strconv"

	"github.com/M-Mahdi-ameri/time_line/internal/domain"
	"github.com/gofiber/fiber/v2"
)

func GetuserByID(repo domain.UserRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi(c.Params("id"))
		user, err := repo.GetByIDu(context.Background(), uint(id))

		if err != nil {
			return c.Status(404).JSON(fiber.Map{
				"error": "user not found",
			})
		}
		return c.JSON(user)
	}
}

func Deleteuser(repo domain.UserRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi(c.Params("id"))
		if err := repo.Deleteu(context.Background(), uint(id)); err != nil {
			return c.Status(404).JSON(fiber.Map{
				"error": "user not found",
			})
		}
		return c.JSON(fiber.Map{
			"message": "user deleted",
		})

	}
}
