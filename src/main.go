package main

import (
	"fmt"
	"gofiber-example/src/adapter/database/sql"
	"gofiber-example/src/adapter/rest"
)

// @title                       API Rest em Fiber
// @contact.name                API Support
// @contact.url                 http://www.swagger.io/support
// @contact.email               support@swagger.io
// @license.name                Apache 2.0
// @license.url                 http://www.apache.org/licenses/LICENSE-2.0.html
// @termsOfService              http://swagger.io/terms/
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
func main() {
	fmt.Println("Inicializando API...")
	sql.DoMigration()
	rest.InitRest()
}
