package router

import (
	"app/auth"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {

	e := echo.New()

	e.POST("/register", auth.Register)
	e.POST("/login", auth.Login)

	r := e.Group("/auth")
	r.Use(echojwt.WithConfig(auth.JwtConfig))

	r.GET("/user", auth.GetAuthUser)

	return e

}
