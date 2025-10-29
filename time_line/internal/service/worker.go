package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/M-Mahdi-ameri/time_line/internal/domain"
	"github.com/go-redis/redis/v8"
)

func StartFanoutWorker(rdb *redis.Client, postRepo domain.PostRepository, userRepo domain.UserRepository, followRepo domain.FollowerRepository) {
	ctx := context.Background()

	log.Println("fan_out worker started and waiting for posts...")

	for {
		res, err := rdb.BRPop(ctx, 0*time.Second, "posts_queue").Result()
		if err != nil {
			log.Printf("redis BRPop error: %s\n", err)
		}

		postIDStr := res[1]

		log.Printf("new post recived from queue: %s\n", postIDStr)

		var postID uint
		fmt.Sscanf(postIDStr, "%d", &postID)

		post, err := postRepo.GetByIDp(ctx, postID)

		if err != nil {
			log.Printf("faild to get follower: %v\n", err)
			continue
		}

		followers, err := followRepo.GetFollowers(ctx, post.AuthorID)
		if err != nil {
			log.Printf("faild to get followers: %v\n", err)
			continue
		}

		for _, followerID := range followers {
			key := fmt.Sprintf("timeline:%d", followerID)
			score := float64(post.CreatedAt.UnixMilli())

			_, err := rdb.ZAdd(ctx, key, &redis.Z{
				Score:  score,
				Member: postID,
			}).Result()

			if err != nil {
				log.Printf("faild to and post %d to follower %d timeline: %v\n", postID, followerID, err)
			}
		}
		log.Printf("post %d disributed to %d followers.\n", postID, len(followers))

	}
}
