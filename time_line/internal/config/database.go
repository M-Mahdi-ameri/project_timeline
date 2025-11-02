package config

import (
	"fmt"
	"log"
	"os"

	"github.com/M-Mahdi-ameri/time_line/internal/domain"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Initmysql() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found , using enviroment variebbles", err)
	}

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("falid to connect to mysql %v", err)
	}

	sqlDb, err := db.DB()
	if err != nil {
		log.Fatalf("fild to get db instance: %v", err)
	}

	if err := sqlDb.Ping(); err != nil {
		log.Fatalf("mysql ping faild :%v", err)
	}
	DB = db
	log.Println("connect to mysql succesfully")

	err = db.AutoMigrate(&domain.User{}, &domain.Post{}, &domain.Follower{})
	if err != nil {
		log.Fatalf("automiragrate faild : %v", err)
	}

	log.Println("database schema miragrated succesfully")

}
