package user

import (
	"errors"
	"gofiber-example/src/internal/model/dbmodel"
	"gofiber-example/src/internal/model/dto"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func setTestEnv() {
	os.Setenv("ENV_VAR_TEST", "test")
}

func TestRegisterUser_OK(t *testing.T) {
	setTestEnv()

	GetObjectFromPostRequest_Mock := func(c *fiber.Ctx, obj interface{}) error {
		user, ok := obj.(*dbmodel.User)
		if !ok {
			return errors.New("Parsing Error")
		}

		user.Name = "Eduardo Andrade"
		user.Email = "eduardo@andrade.com"
		user.Password = "senha_nova"

		return nil
	}

	ReturnOk_Mock := func(c *fiber.Ctx, message string, obj ...interface{}) error {
		assert.Equal(t, "User registered", message)
		return nil
	}

	ReturnInternalServerError_Mock := func(c *fiber.Ctx, message string, err error) error {
		assert.Fail(t, message)
		return nil
	}

	ReturnBadRequest_Mock := func(c *fiber.Ctx, message string, err error) error {
		assert.Fail(t, message)
		return nil
	}
	Create_Mock := func(obj *dbmodel.User) error {
		assert.Equal(t, "Eduardo Andrade", obj.Name)
		assert.Equal(t, "eduardo@andrade.com", obj.Email)
		assert.NotEqual(t, []byte("senha_nova"), obj.Password)
		assert.Nil(t, bcrypt.CompareHashAndPassword([]byte(obj.Password), []byte("senha_nova")))

		return nil
	}

	RegisterUser(
		ReturnOk_Mock,
		GetObjectFromPostRequest_Mock,
		ReturnInternalServerError_Mock,
		ReturnBadRequest_Mock,
		Create_Mock,
	)(nil)
}

func TestRegisterUser_EmptyFields_NOK(t *testing.T) {
	setTestEnv()

	GetObjectFromPostRequest_Mock := func(c *fiber.Ctx, obj interface{}) error {
		_, ok := obj.(*dbmodel.User)
		if !ok {
			return errors.New("Parsing Error")
		}

		return nil
	}

	ReturnOk_Mock := func(c *fiber.Ctx, message string, obj ...interface{}) error {
		assert.Fail(t, message)
		return nil
	}

	ReturnInternalServerError_Mock := func(c *fiber.Ctx, message string, err error) error {
		assert.Fail(t, message)
		return nil
	}

	ReturnBadRequest_Mock := func(c *fiber.Ctx, message string, err error) error {
		assert.Equal(t, "User is not valid", message)
		assert.NotNil(t, err)
		return nil
	}
	Create_Mock := func(obj *dbmodel.User) error {
		assert.Fail(t, "Trying to create a object, but is expected to Fail before")
		return nil
	}

	RegisterUser(
		ReturnOk_Mock,
		GetObjectFromPostRequest_Mock,
		ReturnInternalServerError_Mock,
		ReturnBadRequest_Mock,
		Create_Mock,
	)(nil)
}

func TestLogin_OK(t *testing.T) {
	setTestEnv()

	GetObjectFromPostRequest_Mock := func(c *fiber.Ctx, obj interface{}) error {
		loginRequest, ok := obj.(*dto.LoginRequest)
		if !ok {
			return errors.New("Parsing Error")
		}
		loginRequest.Email = "eduardo@andrade.com"
		loginRequest.Password = "password"
		return nil
	}

	ReturnOk_Mock := func(c *fiber.Ctx, message string, objs ...interface{}) error {
		assert.Equal(t, "Logged in", message)
		for _, obj := range objs {
			loginResponse, ok := obj.(*dto.LoginResponse)
			if !ok {
				return errors.New("Parsing Error")
			}
			assert.Equal(t, "eduardo@andrade.com", loginResponse.Email)
		}
		return nil
	}

	ReturnInternalServerError_Mock := func(c *fiber.Ctx, message string, err error) error {
		assert.Fail(t, message)
		return nil
	}

	ReturnBadRequest_Mock := func(c *fiber.Ctx, message string, err error) error {
		assert.Fail(t, message)
		return nil
	}

	GetUserByEmail_Mock := func(email string) (*dbmodel.User, error) {
		var user dbmodel.User
		user.Password = "$2a$14$WQyku0nWycpL8DanqPjGU.J5/sGKy39zrRPEYdoULiX7tcPSNC7y."
		user.ID = 55
		user.Email = email
		return &user, nil
	}

	AddCookie_Mock := func(c *fiber.Ctx, token string) {
	}

	Login(
		ReturnOk_Mock,
		GetObjectFromPostRequest_Mock,
		ReturnInternalServerError_Mock,
		ReturnBadRequest_Mock,
		GetUserByEmail_Mock,
		AddCookie_Mock,
	)(nil)
}

func TestLogin_Password_Fail(t *testing.T) {
	setTestEnv()

	GetObjectFromPostRequest_Mock := func(c *fiber.Ctx, obj interface{}) error {
		loginRequest, ok := obj.(*dto.LoginRequest)
		if !ok {
			return errors.New("Parsing Error")
		}
		loginRequest.Email = "eduardo@andrade.com"
		loginRequest.Password = "passwordwrong"
		return nil
	}

	ReturnOk_Mock := func(c *fiber.Ctx, message string, objs ...interface{}) error {
		assert.Fail(t, message)
		return nil
	}

	ReturnInternalServerError_Mock := func(c *fiber.Ctx, message string, err error) error {
		assert.Fail(t, message)
		return nil
	}

	ReturnBadRequest_Mock := func(c *fiber.Ctx, message string, err error) error {
		assert.Equal(t, "Review your email/password input", message)
		assert.NotNil(t, err)
		return nil
	}

	GetUserByEmail_Mock := func(email string) (*dbmodel.User, error) {
		var user dbmodel.User
		user.Password = "$2a$14$WQyku0nWycpL8DanqPjGU.J5/sGKy39zrRPEYdoULiX7tcPSNC7y."
		user.ID = 55
		user.Email = email
		return &user, nil
	}

	AddCookie_Mock := func(c *fiber.Ctx, token string) {
		assert.Fail(t, "Trying to add cookie, but is expected to Fail before")
	}

	Login(
		ReturnOk_Mock,
		GetObjectFromPostRequest_Mock,
		ReturnInternalServerError_Mock,
		ReturnBadRequest_Mock,
		GetUserByEmail_Mock,
		AddCookie_Mock,
	)(nil)
}

func TestLogin_Email_Fail(t *testing.T) {
	setTestEnv()

	GetObjectFromPostRequest_Mock := func(c *fiber.Ctx, obj interface{}) error {
		loginRequest, ok := obj.(*dto.LoginRequest)
		if !ok {
			return errors.New("Parsing Error")
		}
		loginRequest.Email = "eduardofake@andrade.com"
		loginRequest.Password = "password"
		return nil
	}

	ReturnOk_Mock := func(c *fiber.Ctx, message string, objs ...interface{}) error {
		assert.Fail(t, message)
		return nil
	}

	ReturnInternalServerError_Mock := func(c *fiber.Ctx, message string, err error) error {
		assert.Equal(t, "User not found", message)
		assert.NotNil(t, err)
		return nil
	}

	ReturnBadRequest_Mock := func(c *fiber.Ctx, message string, err error) error {
		assert.Fail(t, message)
		return nil
	}

	GetUserByEmail_Mock := func(email string) (*dbmodel.User, error) {
		return nil, errors.New("record not found")
	}

	AddCookie_Mock := func(c *fiber.Ctx, token string) {
		assert.Fail(t, "Trying to add cookie, but is expected to Fail before")
	}

	Login(
		ReturnOk_Mock,
		GetObjectFromPostRequest_Mock,
		ReturnInternalServerError_Mock,
		ReturnBadRequest_Mock,
		GetUserByEmail_Mock,
		AddCookie_Mock,
	)(nil)
}
