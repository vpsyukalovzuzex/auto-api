package controllers

import (
	"Demo/db"
	"Demo/models"
	"Demo/usecases"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type IUserController interface {
	SignUp(c echo.Context) error
	SingIn(c echo.Context) error
	Refresh(c echo.Context) error
}

type userController struct {
	uu usecases.IUserUsecase
	ac db.IAuthCache
}

func InitUserController(uu usecases.IUserUsecase, ac db.IAuthCache) IUserController {
	return &userController{uu, ac}
}

func (uc *userController) SignUp(c echo.Context) error {
	userAuth := models.UserAuth{}
	if e := c.Bind(&userAuth); e != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": e.Error(),
		})
	}
	if e := uc.uu.SignUp(userAuth); e != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": e.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Success",
	})
}

func (uc *userController) SingIn(c echo.Context) error {
	userAuth := models.UserAuth{}
	if e := c.Bind(&userAuth); e != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": e.Error(),
		})
	}
	userStored := models.UserStored{}
	if e := uc.uu.SignIn(userAuth, &userStored); e != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": e.Error(),
		})
	}
	ac := models.AuthCredentials{}
	if e := generateAuthCredentials(userStored.Id, &ac); e != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": e.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"data": map[string]string{
			"access_token":  ac.AccessToken,
			"refresh_token": ac.RefreshToken,
		},
	})
}

func (uc *userController) Refresh(c echo.Context) error {
	type body struct {
		RefreshToken string `json:"refresh_token" validate:"required"`
	}
	b := body{}
	if e := c.Bind(&b); e != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": e.Error(),
		})
	}
	secret := os.Getenv("JWTSECRET")
	t, e := jwt.Parse(b.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}
		return nil, errors.New("unexpected signing method")
	})
	if e != nil {
		return echo.ErrUnauthorized
	}
	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		id := int(claims["sub"].(float64))
		userStored := models.UserStored{}
		if e := uc.uu.GetUserById(id, &userStored); e != nil {
			return echo.ErrUnauthorized
		}
		if id != userStored.Id {
			return echo.ErrUnauthorized
		}
		ac := &models.AuthCredentials{}
		if e := generateAuthCredentials(id, ac); e != nil {
			return echo.ErrUnauthorized
		}
		return c.JSON(http.StatusOK, echo.Map{
			"data": map[string]string{
				"access_token":  ac.AccessToken,
				"refresh_token": ac.RefreshToken,
			},
		})
	}
	return echo.ErrUnauthorized
}

func generateAuthCredentials(sub int, ac *models.AuthCredentials) error {
	secret := []byte(os.Getenv("JWTSECRET"))
	at := jwt.New(jwt.SigningMethodHS256)
	atTime := time.Now().Add(time.Minute * 60).Unix()
	atClaims := at.Claims.(jwt.MapClaims)
	atClaims["sub"] = sub
	atClaims["exp"] = atTime
	accessToken, e := at.SignedString(secret)
	if e != nil {
		return e
	}
	rt := jwt.New(jwt.SigningMethodHS256)
	rtTime := time.Now().Add(time.Hour * 8).Unix()
	rtClaims := rt.Claims.(jwt.MapClaims)
	rtClaims["sub"] = sub
	rtClaims["exp"] = rtTime
	refreshToken, e := rt.SignedString(secret)
	if e != nil {
		return e
	}
	ac.AccessToken = accessToken
	ac.RefreshToken = refreshToken
	ac.AccessTokenExp = atTime
	ac.RefreshTokenExp = rtTime
	return nil
}
