package handlers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/M-Mahdi-ameri/time_line/internal/config"
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
		if err := config.DB.Where("user_id = ? OR follower_id = ?", uint(id), uint(id)).Delete(&domain.Follower{}).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "faild to delete followers",
			})
		}
		if err := config.DB.Where("author_id = ?", uint(id)).Delete(&domain.Post{}).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "faild to delete posts",
			})
		}
		if err := repo.Deleteu(context.Background(), uint(id)); err != nil {
			return c.Status(404).JSON(fiber.Map{
				"error": "user not found",
			})
		}
		config.RDB.Del(context.Background(), fmt.Sprintf("timeline:%d", uint(id)))
		keys, _ := config.RDB.Keys(context.Background(), "timeline:*").Result()
		for _, key := range keys {
			config.RDB.ZRem(context.Background(), key, uint(id))
		}

		return c.JSON(fiber.Map{
			"message": "user deleted",
		})

	}
}
