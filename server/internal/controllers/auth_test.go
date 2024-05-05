package controllers

import (
	// "os"
	"testing"
	// "my_project/internal/config"
	// "gorm.io/driver/postgres"
	// "gorm.io/gorm"
)

func TestExample(t *testing.T) {

	// os.Setenv("POSTGRES_HOSTNAME", "localhost")
	// os.Setenv("POSTGRES_USER", "testuser")
	// os.Setenv("POSTGRES_PASSWORD", "testpassword")
	// os.Setenv("POSTGRES_DB", "testdb")
	// os.Setenv("POSTGRES_PORT", "5432")

	// dsn := "host=" + os.Getenv("POSTGRES_HOSTNAME") +
	// 	" user=" + os.Getenv("POSTGRES_USER") +
	// 	" password=" + os.Getenv("POSTGRES_PASSWORD") +
	// 	" dbname=" + os.Getenv("POSTGRES_DB") +
	// 	" port=" + os.Getenv("POSTGRES_PORT") + " sslmode=disable"

	// config.ConnectDatabase()

	// // Attempt to open a test database connection
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err != nil {
	//     t.Fatalf("Failed to connect to test database: %v", err)
	// }

	// // Verify the connection by pinging the database
	// sqlDB, err := db.DB()
	// if err != nil {
	//     t.Fatalf("Failed to get database handle: %v", err)
	// }
	// if err := sqlDB.Ping(); err != nil {
	//     t.Fatalf("Failed to ping database: %v", err)
	// }

	// t.Logf("Successfully connected to test database")
	t.Logf("Running a test")

}

// package controllers

// import (
// 	"bytes"
// 	"encoding/json"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/stretchr/testify/assert"

// 	"my_project/internal/config"
// )

// func setup() *fiber.App {
// 	app := fiber.New()
// 	config.ConnectDatabase()
// 	app.Post("/register", Register)
// 	app.Post("/login", Login)
// 	return app
// }

// func TestRegister(t *testing.T) {
// 	app := setup()

// 	payload := map[string]string{
// 		"username": "testuser",
// 		"email":    "test@example.com",
// 		"password": "password123",
// 	}

// 	body, _ := json.Marshal(payload)
// 	req := httptest.NewRequest("POST", "/register", bytes.NewBuffer(body))
// 	req.Header.Set("Content-Type", "application/json")
// 	resp, _ := app.Test(req)

// 	assert.Equal(t, 201, resp.StatusCode)
// }

// func TestLogin(t *testing.T) {
// 	app := setup()

// 	// Register user first
// 	TestRegister(t)

// 	payload := map[string]string{
// 		"username": "testuser",
// 		"password": "password123",
// 	}

// 	body, _ := json.Marshal(payload)
// 	req := httptest.NewRequest("POST", "/login", bytes.NewBuffer(body))
// 	req.Header.Set("Content-Type", "application/json")
// 	resp, _ := app.Test(req)

// 	assert.Equal(t, 200, resp.StatusCode)
// }
