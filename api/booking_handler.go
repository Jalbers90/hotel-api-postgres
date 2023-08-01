package api

import (
	"strconv"

	"github.com/Jalbers90/hotel-api-postgres/db"
	"github.com/Jalbers90/hotel-api-postgres/types"
	"github.com/gofiber/fiber/v2"
)

type BookingHandler struct {
	store db.BookingStore
}

func NewBookingHandler(store db.BookingStore) *BookingHandler {
	return &BookingHandler{
		store: store,
	}
}

func (h *BookingHandler) HandleGetUserBookings(c *fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Params("userID"))
	if err != nil {
		return err
	}
	bookings, err := h.store.GetUserBookings(c.Context(), int64(userID))
	if err != nil {
		return err
	}
	return c.JSON(bookings)
}

func (h *BookingHandler) HandleInsertBooking(c *fiber.Ctx) error {
	var booking types.Booking
	if err := c.BodyParser(&booking); err != nil {
		return err
	}
	// VALIDATE PAYLOAD
	newBooking, err := h.store.InsertBooking(c.Context(), &booking)
	if err != nil {
		return err
	}
	return c.JSON(newBooking)
}

func (h *BookingHandler) HandleCancelBooking(c *fiber.Ctx) error {
	bookingID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	booking, err := h.store.CancelBooking(c.Context(), int64(bookingID))
	if err != nil {
		return err
	}

	return c.JSON(booking)
}
