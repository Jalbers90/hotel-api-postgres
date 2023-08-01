package api

import (
	"strconv"

	"github.com/Jalbers90/hotel-api-postgres/db"
	"github.com/Jalbers90/hotel-api-postgres/types"
	"github.com/gofiber/fiber/v2"
)

// handle routes, keep decoupled from db stores

type UserHandler struct {
	store db.UserStore
}

func NewUserHandler(store db.UserStore) *UserHandler {
	return &UserHandler{
		store: store,
	}
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	users, err := h.store.GetUsers(c.Context(), map[string]any{})
	if err != nil {
		return err
	}
	return c.JSON(users)
}

func (h *UserHandler) HandleGetUserByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	user, err := h.store.GetUserByID(c.Context(), id)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func (h *UserHandler) HandleInsertUser(c *fiber.Ctx) error {
	var body types.CreateUserParams
	if err := c.BodyParser(&body); err != nil {
		return err
	}
	// VALIDATE LOGIC GO HERE
	// VALIDATED
	user, err := types.CreateUserFromParams(body)
	if err != nil {
		return err
	}
	// fmt.Printf("%+v\n", user)
	insertedUser, err := h.store.InsertUser(c.Context(), user)
	if err != nil {
		return err
	}
	return c.JSON(insertedUser)
}

// func (h *UserHandler) HandleUpdateUser(c *fiber.Ctx) error {
// 	var params types.CreateUserParams
// 	if err := c.BodyParser(&params); err != nil {
// 		return err
// 	}
// }
