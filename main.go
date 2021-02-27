package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "")
	})
	e.GET(
		"/users/:id",
		getUser,
	)
	e.Logger.Fatal(e.Start(":4000"))

}
func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
