package main

import (
	"crud/models"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	/* ENV */
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Use the environment variable for the port
	serverPort := os.Getenv("PORT")
	if serverPort == "" {
		serverPort = "8080" // Default port if not specified
	}

	// Read database configuration from .env file
	host := os.Getenv("DB_HOST")
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT")) // Convert port to int
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	/* Connect DB */
	// Configure PostgreSQL database details here
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, dbPort, user, password, dbname)

	// New logger for detailed SQL logging
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger, // add Logger
	})

	if err != nil {
		panic("failed to connect to database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.Product{})
	fmt.Println("Database migration completed!")

	/* Start server */
	app := fiber.New()
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Go is running.")
	})

	// CRUD routes
	productApi := app.Group("/product")
	productApi.Get("/", func(c *fiber.Ctx) error {
		return models.GetProduct(db, c)
	})
	productApi.Get("/:id", func(c *fiber.Ctx) error {
		return models.GetProductByID(db, c)
	})
	productApi.Post("/", func(c *fiber.Ctx) error {
		return models.CreateProduct(db, c)
	})
	productApi.Put("/:id", func(c *fiber.Ctx) error {
		return models.UpdateProduct(db, c)
	})
	productApi.Delete("/:id", func(c *fiber.Ctx) error {
		return models.DeleteProduct(db, c)
	})

	log.Fatal(app.Listen(":" + serverPort))
}
