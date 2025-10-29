package main

import (
	"log"

	"github.com/M-Mahdi-ameri/time_line/internal/config"
	"github.com/M-Mahdi-ameri/time_line/internal/handlers"
	"github.com/M-Mahdi-ameri/time_line/internal/repository"
	"github.com/M-Mahdi-ameri/time_line/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("ni .env file found ,using envierment variebles")
	}

	config.Initmysql()
	config.Initredis()

	app := fiber.New()

	postRepo := repository.NewGormPostRepo(config.DB)
	userRepo := repository.NewGormUserRepo(config.DB)
	followerRepo := repository.NewGormFollowerRepo(config.DB)

	go service.StartFanoutWorker(config.RDB, postRepo, userRepo, followerRepo)

	handlers.SetupRoutes(app, postRepo, userRepo, followerRepo)

	log.Println("server running on :8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
