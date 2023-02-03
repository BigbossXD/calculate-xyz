package routes

import (
	"github.com/BigbossXD/calculate-xyz/controllers"
	"github.com/BigbossXD/calculate-xyz/services"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, apiV1Prefix string) {

	calculateService := services.NewCalculateService()
	calculateController := controllers.NewCalculateController(calculateService)

	g := e.Group(apiV1Prefix)

	g.GET("/fixed", calculateController.Fixed)
	g.GET("/random", calculateController.BasicRandomInt)
	g.GET("/random-sort", calculateController.BasicRandomIntSort)
	g.GET("/random-odd", calculateController.RandomOdd)
	g.GET("/random-odd-sort", calculateController.RandomOddSort)
	g.GET("/random-even", calculateController.RandomEven)
	g.GET("/random-even-sort", calculateController.RandomEvenSort)
	g.GET("/random-negative", calculateController.RandomNegativeInt)
	g.GET("/random-negative-odd", calculateController.RandomNegativeOddInt)
	g.GET("/random-negative-even", calculateController.RandomNegativeEvenInt)

}
