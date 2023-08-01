package fixtures

import (
	"context"
	"log"
	"time"

	"github.com/Jalbers90/hotel-api-postgres/db"
	"github.com/Jalbers90/hotel-api-postgres/types"
)

func AddUser(store *db.PGUserStore, firstName, lastName, email, password string, admin bool) *types.User {
	params := types.CreateUserParams{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
	}
	user, err := types.CreateUserFromParams(params)
	if err != nil {
		log.Fatal(err)
	}
	user.Admin = admin
	newUser, err := store.InsertUser(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
	return newUser
}

func AddHotel(store *db.PGHotelStore, name, location string, rating uint8) *types.Hotel {
	hotel := types.Hotel{
		Name:     name,
		Location: location,
		Rating:   rating,
	}
	newHotel, err := store.InsertHotel(context.Background(), &hotel)
	if err != nil {
		log.Fatal(err)
	}
	return newHotel
}

func AddRoom(store *db.PGRoomStore, hotelID int64, kind string, price float64) *types.Room {
	room := types.Room{
		HotelID: hotelID,
		Kind:    kind,
		Price:   price,
	}
	newRoom, err := store.InsertRoom(context.Background(), &room)
	if err != nil {
		log.Fatal(err)
	}
	return newRoom
}

func AddBooking(store *db.PGBookingStore, userID, roomID int64, checkin, checkout time.Time, persons int) *types.Booking {
	booking := types.Booking{
		UserID:   userID,
		RoomID:   roomID,
		CheckIn:  checkin,
		CheckOut: checkout,
		Persons:  persons,
	}
	newBooking, err := store.InsertBooking(context.Background(), &booking)
	if err != nil {
		log.Fatal(err)
	}
	return newBooking
}
