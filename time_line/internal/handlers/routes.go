package handlers

import (
	"github.com/M-Mahdi-ameri/time_line/internal/domain"
	"github.com/M-Mahdi-ameri/time_line/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, postRepo domain.PostRepository, userRepo domain.UserRepository, followerRepo domain.FollowerRepository) {

	app.Post("/posts", utils.JWTProtected(), CreatePost(postRepo))
	app.Get("/posts/:id", utils.JWTProtected(), GetpostByID(postRepo))
	app.Get("/posts/author/:id", utils.JWTProtected(), GetPostsByAuthor(postRepo))
	app.Delete("/posts/:id", utils.JWTProtected(), DeletePost(postRepo, followerRepo))

	app.Get("/timeline", utils.JWTProtected(), GetTimeline)

	app.Get("/users/:id", utils.JWTProtected(), GetuserByID(userRepo))
	app.Delete("users/:id", utils.JWTProtected(), Deleteuser(userRepo))

	app.Post("/follow", utils.JWTProtected(), FollowUser(followerRepo))
	app.Post("/unfollow", utils.JWTProtected(), UnfollowUser(followerRepo))
	app.Get("followers/:id", utils.JWTProtected(), GetFollowers(followerRepo))
	app.Get("following/:id", utils.JWTProtected(), GetFollowing(followerRepo))
}
