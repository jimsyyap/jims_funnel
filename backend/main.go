package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"-"` // Ensures password is not serialized
}

var DB *gorm.DB

func initDatabase() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Database connection parameters
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database. \nError: %v", err)
	}

	// Auto Migrate
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatalf("Database migration failed. \nError: %v", err)
	}

	DB = db
	log.Println("Database connection successful")
}

func setupRoutes(app *fiber.App) {
	// Public Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to the API",
			"status":  "healthy",
		})
	})

	// Authentication Routes
	auth := app.Group("/auth")
	auth.Post("/register", registerUser)
	auth.Post("/login", loginUser)
}

func registerUser(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	// Basic validation
	if user.Email == "" || user.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email and password are required",
		})
	}

	// TODO: Add password hashing before storing
	result := DB.Create(user)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"user_id": user.ID,
	})
}

func loginUser(c *fiber.Ctx) error {
	// TODO: Implement proper login logic with password verification
	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"message": "Login endpoint not fully implemented",
	})
}

func main() {
	// Initialize database
	initDatabase()

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName: "MyApp",
	})

	// Middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Be more specific in production
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	}))
	app.Use(logger.New())
	app.Use(recover.New())

	// Setup routes
	setupRoutes(app)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatal(app.Listen(":" + port))
}
