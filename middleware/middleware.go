package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/zakiyalmaya/hotel-management/model"
)

func AuthMiddleware(redcl *redis.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {

		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(model.NewHttpResponse(fiber.StatusUnauthorized, "missing Authorization header", nil))
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate the token signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.NewError(http.StatusUnauthorized, "invalid signing method")
			}
			return []byte("hotel-management-secret-key"), nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(model.NewHttpResponse(fiber.StatusUnauthorized, "invalid or expired token", nil))
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(model.NewHttpResponse(fiber.StatusUnauthorized, "invalid token claims", nil))
		}

		username, ok := claims["username"].(string)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(model.NewHttpResponse(fiber.StatusUnauthorized, "invalid username in token claims", nil))
		}
		c.Locals("username", username)

		userID, ok := claims["user_id"].(float64)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(model.NewHttpResponse(fiber.StatusUnauthorized, "invalid user_id in token claims", nil))
		}
		c.Locals("user_id", int(userID))

		// Check the token in Redis cache
		tokenCache, err := redcl.Get(context.Background(), "jwt-token-"+username).Result()
		if err != nil {
			log.Println("Redis error:", err)
			return c.Status(fiber.StatusUnauthorized).JSON(model.NewHttpResponse(fiber.StatusUnauthorized, "invalid or expired token", nil))
		}

		if tokenCache != tokenString {
			return c.Status(fiber.StatusUnauthorized).JSON(model.NewHttpResponse(fiber.StatusUnauthorized, "invalid or expired token", nil))
		}

		return c.Next()
	}
}
