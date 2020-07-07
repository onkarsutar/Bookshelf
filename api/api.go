package api

import (
	"GoStudy/WebFrameworks/Echo/api/modules/book"
	"net/http"

	"github.com/labstack/echo"
)

// Init init api
func  Init(e *echo.Echo)  {

	e.GET("/", rootRoute)
	o := e.Group("/o")
	// TODO: implement JWT
	r := e.Group("/r")
	
	// r.Use(echoMiddleware.JWTWithConfig(echoMiddleware.JWTConfig{
	// 	Claims:     &model.JwtCustomClaims{},
	// 	SigningKey: []byte(model.JwtKey),
	// }))

	// middleware.Init(e, o, r, c)
	book.Init(o,r)
	
}

func rootRoute(c echo.Context)error{
	return c.String(http.StatusOK, "Welcome...")
}