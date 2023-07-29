package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Jalbers90/hotel-api-postgres/api"
	"github.com/Jalbers90/hotel-api-postgres/db"
	"github.com/gofiber/fiber/v2"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	SSLMode  string
}

func main() {
	config := Config{
		Host:     "localhost",
		Port:     5432,
		User:     "admin",
		Password: "admin",
		Database: "hotel",
		SSLMode:  "disable",
	}
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
		config.SSLMode,
	)
	pgdb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(connStr)))
	bunDB := bun.NewDB(pgdb, pgdialect.New())
	defer bunDB.Close()

	app := fiber.New(fiber.Config{})
	apiv1 := app.Group("/api/v1")

	userStore := db.NewPGUserStore(bunDB)
	userHandler := api.NewUserHandler(userStore)
	// handler for each route
	apiv1.Get("/user", userHandler.HandleGetUsers)
	apiv1.Get("/user/:id", userHandler.HandleGetUserByID)

	log.Fatal(app.Listen(":8000"))
}
