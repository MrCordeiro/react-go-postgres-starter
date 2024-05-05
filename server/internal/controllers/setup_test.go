package controllers

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"testing"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"my_project/internal/config"
)

const dbName = "testdb"

var db *gorm.DB

func TestMain(m *testing.M) {
	setupTestDatabase()
	code := m.Run()
	teardownTestDatabase()
	os.Exit(code)
}

/*
Creates a new test database and migrates the schema.
It also assigns the test database instance to the `testDB` variable.
*/
func setupTestDatabase() {
	log.Println("Setting up test database")
	// Connect to the default database
	err := godotenv.Load(filepath.Join(config.ProjectDir, ".env"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := getDSN()
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Create a test database
	db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s;", dbName))
	db.Exec(fmt.Sprintf("CREATE DATABASE %s;", dbName))

	// Connect to the test database
	testDSN := dsn + " dbname=" + dbName
	_, err = gorm.Open(postgres.Open(testDSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to test database: %v", err)
	}

	// Run migrations (Replace `YourModel` with your actual model)
	// testDB.AutoMigrate(&YourModel{}) // Add other models as needed

	// Assign the database instance to be used in tests
}

func getDSN() string {
	dsn := "host=" + os.Getenv("POSTGRES_HOSTNAME") +
		" user=" + os.Getenv("POSTGRES_USER") +
		" password=" + os.Getenv("POSTGRES_PASSWORD") +
		" sslmode=disable"
	return dsn
}

func teardownTestDatabase() {
	log.Println("Tearing down test database")

	// Terminate all connections to the test database
	db.Exec(
		fmt.Sprintf(
			"SELECT pg_terminate_backend(pg_stat_activity.pid) FROM pg_stat_activity "+
				"WHERE pg_stat_activity.datname = '%s' AND pid <> pg_backend_pid();",
			dbName,
		),
	)

	// Drop the test database
	db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s;", dbName))

	// Close the test database connection
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get underlying SQL database: %v", err)
	}
	sqlDB.Close()
}
