package rest

import (
	userdb "gofiber-example/src/adapter/database/sql/user"
	"gofiber-example/src/adapter/rest/v1/controller/user"
	controller_helper "gofiber-example/src/helper/controller"
	controller_error "gofiber-example/src/helper/controller/error"
	controller_success "gofiber-example/src/helper/controller/success"
	"gofiber-example/src/helper/jwt"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	_ "gofiber-example/src/docs"
)

func InitRest() {
	app := fiber.New()

	app.Use(cors.New())

	app.Get("/swagger/*", swagger.Handler)

	v1Group := app.Group("/v1")
	{
		userGroup := v1Group.Group("/user")
		{
			userGroup.Post("/register", user.RegisterUser(
				controller_success.ReturnOk,
				controller_helper.GetObjectFromPostRequest,
				controller_error.ReturnInternalServerError,
				controller_error.ReturnBadRequest,
				userdb.Create,
			))
			userGroup.Post("/login", user.Login(
				controller_success.ReturnOk,
				controller_helper.GetObjectFromPostRequest,
				controller_error.ReturnInternalServerError,
				controller_error.ReturnBadRequest,
				userdb.GetUserByEmail,
				jwt.AddCookie,
			))
			userGroup.Get("/", user.User)
			userGroup.Post("/logout", user.Logout)
		}
	}

	err := app.Listen(":5000")

	if err != nil {
		panic(err.Error())
	}
}
