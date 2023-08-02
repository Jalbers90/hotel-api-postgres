package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Jalbers90/hotel-api-postgres/db"
	"github.com/Jalbers90/hotel-api-postgres/db/fixtures"
	"github.com/Jalbers90/hotel-api-postgres/types"
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

const (
	DAY = time.Second * 86400
)

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

	// RESET TABLES
	err := bunDB.ResetModel(context.Background(), (*types.User)(nil), (*types.Hotel)(nil), (*types.Room)(nil), (*types.Booking)(nil))
	if err != nil {
		log.Fatal(err)
	}
	userStore := db.NewPGUserStore(bunDB)
	hotelStore := db.NewPGHotelStore(bunDB)
	roomStore := db.NewPGRoomStore(bunDB)
	bookingStore := db.NewPGBookingStore(bunDB)
	// _ = bookingStore
	user1 := fixtures.AddUser(userStore, "John", "Albers", "john@albers.com", "password", true)
	user2 := fixtures.AddUser(userStore, "Amy", "Albers", "amy@albers.com", "password", false)

	hotel1 := fixtures.AddHotel(hotelStore, "Hilton", "Texas", 5)

	room1 := fixtures.AddRoom(roomStore, hotel1.HotelID, "King", 99.99)
	//_ = room1
	room2 := fixtures.AddRoom(roomStore, hotel1.HotelID, "Queen", 109.99)
	fixtures.AddRoom(roomStore, hotel1.HotelID, "Suite", 199.99)

	fixtures.AddBooking(bookingStore, user1.ID, room1.RoomID, time.Now(), time.Now().Add(DAY), 1)
	fixtures.AddBooking(bookingStore, user2.ID, room2.RoomID, time.Now(), time.Now().Add(DAY), 2)

	hotel2 := fixtures.AddHotel(hotelStore, "marriot", "Florida", 4)
	fixtures.AddRoom(roomStore, hotel2.HotelID, "King", 99.99)
	fixtures.AddRoom(roomStore, hotel2.HotelID, "Queen", 109.99)
	fixtures.AddRoom(roomStore, hotel2.HotelID, "Suite", 199.99)
	hotel3 := fixtures.AddHotel(hotelStore, "Hotel Motel Holiday Inn", "Chiraq", 1)
	fixtures.AddRoom(roomStore, hotel3.HotelID, "King", 99.99)
	fixtures.AddRoom(roomStore, hotel3.HotelID, "Queen", 109.99)
	fixtures.AddRoom(roomStore, hotel3.HotelID, "Suite", 199.99)
	fmt.Printf("%+v\n", user1)
	fmt.Printf("%+v\n", hotel3)
}
