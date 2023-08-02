package api

import (
	"fmt"
	"os"
	"time"

	"github.com/Jalbers90/hotel-api-postgres/db"
	"github.com/Jalbers90/hotel-api-postgres/types"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthentication(userStore db.UserStore) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// fmt.Println("--- JWT Auth ---")
		token, ok := c.GetReqHeaders()["X-Api-Token"]
		if !ok {
			return fmt.Errorf("unauthorized")
		}
		claims, err := validateToken(token)
		if err != nil {
			return err
		}
		// check token expiration
		expires := claims["expires"].(float64)
		if time.Now().Unix() > int64(expires) {
			return fmt.Errorf("unauthorized")
		}
		// user, err := userStore.GetUserByID(c.Context(), userID)
		if err != nil {
			return fmt.Errorf("unauthorized")
		}
		// set the current authenticated user to the context
		c.Context().SetUserValue("user", types.Map{"userID": claims["id"], "email": claims["email"], "expires": claims["expires"]})
		return c.Next()
	}
}

func validateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("Unexpected signing method:", token.Header["alg"])
			return nil, fmt.Errorf("unauthorized")
		}
		secret := os.Getenv("JWT_SECRET")
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secret), nil
	})
	if err != nil {
		fmt.Println("failed to parse token", err)
		return nil, fmt.Errorf("unauthorized")
	}
	if !token.Valid {
		fmt.Println("invalid token")
		return nil, fmt.Errorf("unauthorized")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("unauthorized")
	}
	return claims, nil
}
