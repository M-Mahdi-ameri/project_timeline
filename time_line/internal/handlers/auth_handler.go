package handlers

import (
	"time"

	"github.com/M-Mahdi-ameri/time_line/internal/config"
	"github.com/M-Mahdi-ameri/time_line/internal/domain"
	"github.com/M-Mahdi-ameri/time_line/internal/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type registerRequest struct {
	Username string `json:"username" valid:"required"`
	Email    string `json:"email" valid:"email,required"`
	Password string `json:"password" valid:"required,length(6|32)"`
}

type loginRequest struct {
	Email    string `json:"email" valid:"required,email"`
	Password string `json:"password" valid:"required"`
}

func Register(c *fiber.Ctx) error {
	var req registerRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	var existing domain.User

	if err := config.DB.Where("email: ?", req.Email).First(existing).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	user := domain.User{
		Username:  req.Username,
		Email:     req.Email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
	}
	if err := config.DB.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "faild to create user",
		})
	}
	return c.Status(201).JSON(fiber.Map{
		"message": "User reqisted succesfully",
	})
}

func Login(c *fiber.Ctx) error {
	var req loginRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}
	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	var user domain.User

	if err := config.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid email or password",
		})
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid email or password",
		})
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "faild to generate token",
		})
	}
	return c.Status(500).JSON(fiber.Map{
		"token": token,
	})
}
