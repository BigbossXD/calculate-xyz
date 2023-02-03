package main

import (
	"os"
	"time"

	"github.com/BigbossXD/calculate-xyz/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	time.LoadLocation("Asia/Bangkok")

	godotenv.Load(".env")
	appPort := os.Getenv("APP_PORT")

	e := echo.New()

	// /* CORS */
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	apiV1Prefix := "api/v1"
	routes.InitRoutes(e, apiV1Prefix)
	// routes.InitRoutes(e, apiV1Prefix)
	// routes.InitMachineRoutes(e, apiV1Prefix)
	// routes.InitTransectionRoutes(e, apiV1Prefix)

	e.Logger.Fatal(e.Start(":" + appPort))
}
