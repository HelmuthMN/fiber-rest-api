package jwt

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
)

const SecretKey = "secret"

type GenerateTokenFn func(userID uint) (string, error)

func GenerateToken(userID uint) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: strconv.Itoa(int(userID)),
	})

	return claims.SignedString([]byte(SecretKey))
}

type AddCookieFn func(c *fiber.Ctx, token string)

func AddCookie(c *fiber.Ctx, token string) {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
}
