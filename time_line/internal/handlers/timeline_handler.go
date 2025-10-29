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
	userID := c.Locals("user_id").(uint)

	limitParam := c.Query("limit", "10")
	beforeParam := c.Query("before", "")

	limit, err := strconv.Atoi(limitParam)
	if err != nil {
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
			"error": "faild to load timeline",
		})
	}

	if len(res) == 0 {
		return c.JSON([]domain.Post{})
	}

	postID := make([]uint, 0, len(res))
	for _, z := range res {
		id, _ := strconv.Atoi(z.Member.(string))
		postID = append(postID, uint(id))
	}

	var posts []domain.Post

	if err := config.DB.Where("id IN ?", postID).Order("created_at DESC").Find(&posts).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "faild to fetch posts",
		})
	}
	return c.JSON(posts)
}
