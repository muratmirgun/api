package router

import (
	"github.com/labstack/echo/v4"
	"net/http"

)

func GetUser(c echo.Context) error {
	id := c.Param("id")
	sum := id + "get user"
	return c.String(http.StatusOK, sum)
}

func PutUser(c echo.Context) error {
	id := c.Param("id")
	sum := id + "Put user"
	return c.String(http.StatusOK, sum)
}

func PostUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK,id )
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}