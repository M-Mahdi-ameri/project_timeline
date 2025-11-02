package handlers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/M-Mahdi-ameri/time_line/internal/config"
	"github.com/M-Mahdi-ameri/time_line/internal/domain"
	"github.com/gofiber/fiber/v2"
)

func FollowUser(repo domain.FollowerRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		type followrReq struct {
			UserID uint `json:"user_id"`
		}
		var req followrReq
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "invalid body",
			})
		}
		followerID := c.Locals("user_id").(uint)

		if followerID == req.UserID {
			return c.Status(400).JSON(fiber.Map{
				"error": "you cant follow yourself",
			})
		}
		if err := repo.Follow(context.Background(), req.UserID, followerID); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "faild to follow",
			})
		}
		return c.JSON(fiber.Map{
			"message": "followed seccessfully",
		})
	}
}
func UnfollowUser(repo domain.FollowerRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		type unfollowrReq struct {
			UserID uint `json:"user_id"`
		}
		var req unfollowrReq
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "invalid body",
			})
		}
		followerID := c.Locals("user_id").(uint)
		if err := repo.Unfollow(context.Background(), req.UserID, followerID); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "faild to unfollow",
			})
		}

		var authorPosts []domain.Post
		if err := config.DB.Where("author_id = ?", req.UserID).Find(&authorPosts).Error; err == nil {
			for _, p := range authorPosts {
				config.RDB.ZRem(context.Background(), fmt.Sprintf("timeline:%d", followerID), p.ID)
			}
		}
		return c.Status(200).JSON(fiber.Map{
			"message": "unfollowed seccessfully",
		})
	}
}

func GetFollowers(repo domain.FollowerRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi(c.Params("id"))
		followers, err := repo.GetFollowers(context.Background(), uint(id))

		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "faild to fetch followers",
			})
		}
		return c.JSON(fiber.Map{
			"followers": followers})
	}
}
func GetFollowing(repo domain.FollowerRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi(c.Params("id"))
		following, err := repo.GetFollowing(context.Background(), uint(id))

		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "faild to fetch following",
			})
		}
		return c.JSON(fiber.Map{
			"following": following})
	}
}
