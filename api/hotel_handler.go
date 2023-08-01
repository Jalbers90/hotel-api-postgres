package api

import (
	"strconv"

	"github.com/Jalbers90/hotel-api-postgres/db"
	"github.com/gofiber/fiber/v2"
)

type HotelHandler struct {
	store db.HotelStore
}

func NewHotelHandler(store db.HotelStore) *HotelHandler {
	return &HotelHandler{
		store: store,
	}
}

func (h *HotelHandler) HandleGetHotels(c *fiber.Ctx) error {
	users, err := h.store.GetHotels(c.Context(), map[string]any{})
	if err != nil {
		return err
	}
	return c.JSON(users)
}

func (h *HotelHandler) HandleGetHotelByID(c *fiber.Ctx) error {
	hotelID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	hotel, err := h.store.GetHotelByID(c.Context(), hotelID)
	if err != nil {
		return err
	}
	return c.JSON(hotel)
}

func (h *HotelHandler) HandleGetHotelRooms(c *fiber.Ctx) error {
	hotelID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	rooms, err := h.store.GetHotelRooms(c.Context(), hotelID)
	return c.JSON(rooms)
}
