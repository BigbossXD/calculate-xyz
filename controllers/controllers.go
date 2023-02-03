package controllers

import (
	"net/http"

	"github.com/BigbossXD/calculate-xyz/models/responses"
	"github.com/BigbossXD/calculate-xyz/services"
	"github.com/labstack/echo/v4"
)

type CalculateController struct {
	calculateService *services.CalculateService
}

func NewCalculateController(calculateService *services.CalculateService) *CalculateController {
	return &CalculateController{
		calculateService: calculateService,
	}
}

func BasicResponse(c echo.Context, r *responses.BasicResultResponse, err error) error {
	if err != nil {
		response := responses.ErrorBaseResponse{
			Code:    "00400",
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}
	response := responses.SuccessBaseResponse{
		Code:    "00000",
		Message: "Success",
		Data:    r,
	}
	return c.JSON(http.StatusOK, response)
}

func (h *CalculateController) Fixed(c echo.Context) error {
	result, err := h.calculateService.Fixed()
	return BasicResponse(c, result, err)
}

func (h *CalculateController) BasicRandomInt(c echo.Context) error {
	result, err := h.calculateService.BasicRandomInt()
	return BasicResponse(c, result, err)
}

func (h *CalculateController) BasicRandomIntSort(c echo.Context) error {
	result, err := h.calculateService.BasicRandomIntSort()
	return BasicResponse(c, result, err)
}

func (h *CalculateController) RandomOdd(c echo.Context) error {
	result, err := h.calculateService.RandomOdd()
	return BasicResponse(c, result, err)
}

func (h *CalculateController) RandomOddSort(c echo.Context) error {
	result, err := h.calculateService.RandomOddSort()
	return BasicResponse(c, result, err)
}

func (h *CalculateController) RandomEven(c echo.Context) error {
	result, err := h.calculateService.RandomEven()
	return BasicResponse(c, result, err)
}

func (h *CalculateController) RandomEvenSort(c echo.Context) error {
	result, err := h.calculateService.RandomEvenSort()
	return BasicResponse(c, result, err)
}

func (h *CalculateController) RandomNegativeInt(c echo.Context) error {
	result, err := h.calculateService.RandomNegativeInt()
	return BasicResponse(c, result, err)
}

func (h *CalculateController) RandomNegativeOddInt(c echo.Context) error {
	result, err := h.calculateService.RandomNegativeOddInt()
	return BasicResponse(c, result, err)
}

func (h *CalculateController) RandomNegativeEvenInt(c echo.Context) error {
	result, err := h.calculateService.RandomNegativeEvenInt()
	return BasicResponse(c, result, err)
}
