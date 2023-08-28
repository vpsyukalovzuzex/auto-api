package controllers

import (
	"Demo/models"
	"Demo/usecases"
	"Demo/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ICarController interface {
	GetUserCars(c echo.Context) error
	GetUserEngines(c echo.Context) error
	GetCarEngine(c echo.Context) error
}
type carController struct {
	cu usecases.ICarUsecase
}

func InitCarController(cu usecases.ICarUsecase) ICarController {
	return &carController{cu}
}

func (cc *carController) GetUserCars(c echo.Context) error {
	r := []models.UserCar{}
	sub := utils.GetSub(c)
	_ = cc.cu.GetAllUserCarsByUserId(sub, &r)
	return c.JSON(http.StatusOK, echo.Map{
		"data": r,
	})
}

func (cc *carController) GetUserEngines(c echo.Context) error {
	r := []models.Engine{}
	sub := utils.GetSub(c)
	_ = cc.cu.GetAllUserEnginesByUserId(sub, &r)
	return c.JSON(http.StatusOK, echo.Map{
		"data": r,
	})
}

func (cc *carController) GetCarEngine(c echo.Context) error {
	r := models.Engine{}
	id, e := strconv.Atoi(c.Param("id"))
	if e != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Bad engine id value",
		})
	}
	if e := cc.cu.GetCarEngine(id, &r); e != nil {
		e := e.(*utils.CustomError)
		return c.JSON(e.Code, echo.Map{
			"error": e.Description,
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"data": r,
	})
}
