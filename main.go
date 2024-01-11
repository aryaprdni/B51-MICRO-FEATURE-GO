package main

import (
	"B51-MICRO-FEATURE-GO/database"
	"B51-MICRO-FEATURE-GO/pkg/mysql"
	"B51-MICRO-FEATURE-GO/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	mysql.DatabaseInit()
	database.RunMigration()

	routes.Routes(e.Group("/api/v1"))

	e.Logger.Fatal(e.Start("localhost:5000"))
}
