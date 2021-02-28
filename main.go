package main

import (
	"api/models"
	"api/router"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {

	db := models.InitDB("sqlite.db")

	deletes := router.DeleteUser

	// Router
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "")
	})
	e.GET("/users/:id", router.GetUser(db))
	e.POST("/users/new/:id", router.PostUser(db))
	e.DELETE("/users/delete/:id", deletes)
	//logger and server start
	e.Logger.Fatal(e.Start(":4000"))

}
