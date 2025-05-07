package utils

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

var Token = struct {
	VerifyToken func(string) fiber.Handler
	GenerateToken func(string) (string, error)
}{
	VerifyToken: func(secret string) fiber.Handler {
		return func(c *fiber.Ctx) error {
			tokenString := c.Get("Authorization")
			if tokenString == "" {
				return c.Status(401).SendString("Token manquant")
			}
			
			_, err := VerifyToken(tokenString, secret)
			if err != nil {
				return c.Status(401).SendString("Token invalide")
			}
			
			return c.Next()
		}
	},
	GenerateToken: func(userID string) (string, error) {
		return GenerateTokens(userID, "your-secret-key")
	},
}

func GenerateTokens(userID string, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 72).Unix(), // 1 hour expiration
		"iat": time.Now().Unix(),
		"type": "access",
	})
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString string, secret string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}
	return token.Claims.(jwt.MapClaims)["sub"].(string), nil
}
