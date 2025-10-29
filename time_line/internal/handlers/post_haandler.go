package handlers

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/M-Mahdi-ameri/time_line/internal/config"
	"github.com/M-Mahdi-ameri/time_line/internal/domain"
	"github.com/gofiber/fiber/v2"
)

func CreatePost(repo domain.PostRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var post domain.Post
		if err := c.BodyParser(&post); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "invalid body",
			})
		}
		userID := c.Locals("user_id").(uint)
		post.AuthorID = userID
		if err := repo.Create(context.Background(), &post); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "faild to create post",
			})
		}

		queueKey := "posts_queue"
		postIDStr := fmt.Sprintf("%d", post.ID)
		if err := config.RDB.LPush(context.Background(), queueKey, postIDStr).Err(); err != nil {
			log.Printf("faild to push post %dto redis queue: %v\n", post.ID, err)
		} else {
			log.Printf("post %d pushed to redis queue successfully\n", post.ID)
		}
		return c.Status(201).JSON(post)
	}
}

func GetpostByID(repo domain.PostRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi(c.Params("id"))
		post, err := repo.GetByIDp(context.Background(), uint(id))
		if err != nil {
			return c.Status(404).JSON(fiber.Map{
				"error": "post not found",
			})
		}
		return c.Status(201).JSON(post)
	}
}

func GetPostsByAuthor(repo domain.PostRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi(c.Params("id"))
		posts, err := repo.GetPostsByAuthor(context.Background(), uint(id))
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "fiald to fetch posts",
			})
		}
		return c.Status(201).JSON(posts)
	}
}
func DeletePost(repo domain.PostRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi(c.Params("id"))
		if err := repo.Deletep(context.Background(), uint(id)); err != nil {
			return c.Status(404).JSON(fiber.Map{
				"error": "post not found",
			})
		}
		return c.JSON(fiber.Map{
			"meesage": "post deleted",
		})
	}
}
