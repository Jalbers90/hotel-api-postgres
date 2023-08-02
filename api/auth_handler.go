package api

import (
	"fmt"
	"os"
	"time"

	"github.com/Jalbers90/hotel-api-postgres/db"
	"github.com/Jalbers90/hotel-api-postgres/types"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	store db.UserStore
}

func NewAuthHandler(userStore db.UserStore) *AuthHandler {
	return &AuthHandler{
		store: userStore,
	}
}

type AuthReqBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthSuccessResponse struct {
	User  *types.User
	Token string `json:"token"`
}

func (h *AuthHandler) HandleAuthenticateUser(c *fiber.Ctx) error {
	var authBody AuthReqBody
	if err := c.BodyParser(&authBody); err != nil {
		return err
	}
	// have email and password
	user, err := h.store.GetUserByEmail(c.Context(), authBody.Email)
	if err != nil {
		return fmt.Errorf("unauthorized")
	}
	// check password
	if bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(authBody.Password)) != nil {
		return fmt.Errorf("unauthorized")
	}
	// create token
	token := CreateTokenFromUser(user)
	return c.JSON(AuthSuccessResponse{User: user, Token: token})
}

func CreateTokenFromUser(user *types.User) string {
	now := time.Now()
	expires := now.Add(time.Hour * 24).Unix()
	claims := jwt.MapClaims{
		"id":      user.ID,
		"email":   user.Email,
		"expires": expires,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Println("error signing token:", err)
	}
	return tokenStr
}
