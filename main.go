package main

import (
	"database/sql"
	"os"

	"github.com/anti-duhring/goncurrency/internal/db"
	"github.com/anti-duhring/goncurrency/internal/http"
	"github.com/gofiber/fiber/v2"
)

var (
	App *fiber.App
	DB  *sql.DB
)

func main() {
	var err error

	DB, err = db.Init()
	if err != nil {
		os.Exit(0)
	}

	App = http.Init(DB)

	port := os.Getenv("API_PORT")
	if port == "" {
		port = "3000"
	}

	App.Listen(":" + port)
}
