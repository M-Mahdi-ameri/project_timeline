package handlers

import (
	"context"
	"strconv"

	"github.com/M-Mahdi-ameri/time_line/internal/config"
	"github.com/M-Mahdi-ameri/time_line/internal/domain"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func GetTimeline(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return c.Status(401).JSON(fiber.Map{
			"error": "unauthorizwd:invalid token payload",
		})
	}

	limitParam := c.Query("limit", "10")
	beforeParam := c.Query("before", "")

	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit <= 0 {
		limit = 10
	}

	var maxScore string

	if beforeParam != "" {
		maxScore = "(" + beforeParam
	} else {
		maxScore = "+inf"
	}

	ctx := context.Background()
	key := "timeline:" + strconv.Itoa(int(userID))

	res, err := config.RDB.ZRevRangeByScoreWithScores(ctx, key, &redis.ZRangeBy{
		Max:    maxScore,
		Min:    "-inf",
		Offset: 0,
		Count:  int64(limit),
	}).Result()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "faild to load timeline from redis",
		})
	}

	if len(res) == 0 {
		return c.JSON([]domain.Post{})
	}

	postID := make([]uint, 0, len(res))
	for _, z := range res {
		switch v := z.Member.(type) {
		case string:
			id, err := strconv.Atoi(v)
			if err == nil {
				postID = append(postID, uint(id))
			}
		case int:
			postID = append(postID, uint(v))
		case int64:
			postID = append(postID, uint(v))
		default:
			continue
		}

	}
	if len(postID) == 0 {
		return c.JSON([]domain.Post{})
	}
	var posts []domain.Post

	if err := config.DB.Where("id IN ?", postID).Order("created_at DESC").Find(&posts).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "faild to fetch posts",
		})
	}
	return c.JSON(posts)
}
