package middleware

import (
	jwtHelper "gofiber-example/src/helper/jwt"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
)

func AuthenticateCookie(c *fiber.Ctx) (*jwt.StandardClaims, error) {

	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtHelper.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims := token.Claims.(*jwt.StandardClaims)

	return claims, nil
}

func RemoveCookie(c *fiber.Ctx) {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
}
