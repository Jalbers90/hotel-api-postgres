package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Jalbers90/hotel-api-postgres/api"
	"github.com/Jalbers90/hotel-api-postgres/db"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	config := Config{
		Host:     "localhost",
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("PG_USER"),
		Password: os.Getenv("PG_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
		SSLMode:  "disable",
	}
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
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
	hotelStore := db.NewPGHotelStore(bunDB)
	bookingStore := db.NewPGBookingStore(bunDB)

	// USER ROUTES
	userHandler := api.NewUserHandler(userStore)
	apiv1.Get("/user", userHandler.HandleGetUsers)
	apiv1.Get("/user/:id", userHandler.HandleGetUserByID)
	apiv1.Post("/user", userHandler.HandleInsertUser)

	// HOTEL ROUTES
	hotelHandler := api.NewHotelHandler(hotelStore)
	apiv1.Get("/hotel", hotelHandler.HandleGetHotels)
	apiv1.Get("/hotel/:id", hotelHandler.HandleGetHotelByID)
	apiv1.Get("/hotel/:id/rooms", hotelHandler.HandleGetHotelRooms)

	// BOOKING ROUTES
	bookingHandler := api.NewBookingHandler(bookingStore)
	// apiv1.Get("/booking/hotel/:hotelID", bookingHandler.HandleGetHotelBookings)
	apiv1.Post("/booking/book", bookingHandler.HandleInsertBooking)          // insert new booking, booking payload, must include dates/userid/roomid
	apiv1.Get("/booking/user/:userID", bookingHandler.HandleGetUserBookings) // get all bookings for a user
	apiv1.Get("/booking/:id/cancel", bookingHandler.HandleCancelBooking)
	// apiv1.Post("/booking/room/:roomID", bookingHandler.HandleGetRoomBookings) // get bookings for specific roomID, optional date payload
	// apiv1.Get("/booking/user/:userID", bookingHandler.HandleGetUserBookings)

	log.Fatal(app.Listen(":8000"))
}
