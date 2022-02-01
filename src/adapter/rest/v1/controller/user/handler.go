package user

import (
	"errors"

	"gofiber-example/src/adapter/database/sql"
	userdb "gofiber-example/src/adapter/database/sql/user"
	"gofiber-example/src/adapter/rest/middleware"
	controller_helper "gofiber-example/src/helper/controller"
	controller_error "gofiber-example/src/helper/controller/error"
	controller_success "gofiber-example/src/helper/controller/success"
	jwtHelper "gofiber-example/src/helper/jwt"
	"gofiber-example/src/helper/validate"
	"gofiber-example/src/internal/model/dbmodel"
	"gofiber-example/src/internal/model/dto"

	"github.com/gofiber/fiber/v2"

	"golang.org/x/crypto/bcrypt"
)

// Register User on API
// @Summary      Create a new user
// @Description  Create a new user based on user model
// @Tags         Register User
// @Accept       json
// @Produce      json
// @Param        payload  body      dbmodel.User           true  "User that will be registered"
// @Success      200      {object}  dbmodel.User           "Response if object is found"
// @Success      500      {object}  dbmodel.ErrorResponse  "Response if object was not found"
// @Router       /v1/user/register [post]
func RegisterUser(
	ReturnOk controller_success.ReturnFn,
	GetObjectFromPostRequest controller_helper.GetObjectFromPostRequestFn,
	ReturnInternalServerError controller_error.ReturnFn,
	ReturnBadRequest controller_error.ReturnFn,
	CreateUser userdb.CreateFn,
) func(c *fiber.Ctx) error {

	return func(c *fiber.Ctx) error {
		var user dbmodel.User

		if err := GetObjectFromPostRequest(c, &user); err != nil {
			return ReturnBadRequest(c, "Review your input", err)
		}

		err := validate.ValidateUser(user)
		if err != nil {
			return ReturnBadRequest(c, "User is not valid", err)
		}

		password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
		user.Password = string(password)

		if err := CreateUser(&user); err != nil {
			return ReturnInternalServerError(c, "Could not create user", errors.New("review your input"))
		}
		return ReturnOk(c, "User registered")
	}
}

// Log with a existing user
// @Summary      Log the user
// @Description  Log the user in API
// @Tags         Login
// @Accept       json
// @Produce      json
// @Param        payload  body      dbmodel.User  true  "User"
// @Success      200      {object}  dbmodel.User  "teste"
// @Success      400      {object}  dbmodel.ErrorResponse
// @Success      403      {object}  dbmodel.ErrorResponse
// @Success      500      {object}  dbmodel.ErrorResponse
// @Router       /v1/user/login [post]
func Login(
	ReturnOk controller_success.ReturnFn,
	GetObjectFromPostRequest controller_helper.GetObjectFromPostRequestFn,
	ReturnInternalServerError controller_error.ReturnFn,
	ReturnBadRequest controller_error.ReturnFn,
	GetUserByEmail userdb.GetUserByEmailFn,
	AddCookie jwtHelper.AddCookieFn,
) func(c *fiber.Ctx) error {

	return func(c *fiber.Ctx) error {
		var data dto.LoginRequest

		if err := GetObjectFromPostRequest(c, &data); err != nil {
			return ReturnBadRequest(c, "Review your input", err)
		}

		if data.Email == "" {
			return ReturnBadRequest(c, "Review your email/password input", errors.New("email/password is not valid"))
		}

		if len(data.Password) == 0 {
			return ReturnBadRequest(c, "Review your email/password input", errors.New("email/password is not valid"))
		}

		user, err := GetUserByEmail(data.Email)
		if err != nil {
			return ReturnInternalServerError(c, "User not found", errors.New("email/password is not valid"))
		}

		if user.ID == 0 {
			return ReturnInternalServerError(c, "User not found", errors.New("email/password is not valid"))
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
			return ReturnBadRequest(c, "Review your email/password input", err)
		}

		token, err := jwtHelper.GenerateToken(user.ID)

		if err != nil {
			return ReturnInternalServerError(c, "Could not login", err)
		}

		AddCookie(c, token)

		return ReturnOk(c, "Logged in", &dto.LoginResponse{
			Token:     token,
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Name:      user.Name,
			Email:     user.Email,
		})
	}
}

const SecretKey = "secret"

// User func gets the user
// @Description  Get the actual user.
// @Summary      get the user
// @Tags         User
// @Accept       json
// @Produce      json
// @Success      200  {object}  dbmodel.User  "teste"
// @Router       /v1/user/ [get]
func User(c *fiber.Ctx) error {
	claims, err := middleware.AuthenticateCookie(c)

	if err != nil {
		controller_error.ReturnUnauthorized(c, "Unauthorized", err)
	}

	var user dbmodel.User

	sql.DB.Where("id = ?", claims.Issuer).First(&user)

	return c.JSON(user)
}

// Logout from api server godoc
// @Summary      Logout from server
// @Description  Logout from api server
// @Tags         Logout
// @Accept       json
// @Produce      json
// @Success      200  {object}  dbmodel.User
// @Router       /v1/user/logout [post]
func Logout(c *fiber.Ctx) error {
	middleware.RemoveCookie(c)

	return controller_success.ReturnOk(c, "Logged out successfully")
}
