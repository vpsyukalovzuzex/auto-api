package routers

import (
	"Demo/controllers"
	"os"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Create(uc controllers.IUserController, cc controllers.ICarController) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	jwt := echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("JWTSECRET")),
	})
	e.POST("/signup", uc.SignUp)
	e.POST("/signin", uc.SingIn)
	e.POST("/refresh", uc.Refresh)
	e.GET("/cars", cc.GetUserCars, jwt)
	e.GET("/engines", cc.GetUserEngines, jwt)
	e.GET("/car/:id/engine", cc.GetCarEngine, jwt)
	return e
}
