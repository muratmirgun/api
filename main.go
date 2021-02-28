package main

import (
	"api/router"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	// router imports
    gets := router.GetUser
    posts := router.PostUser
    deletes := router.DeleteUser

    // Router
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "")
	})
	e.GET("/users/:id", gets)
    e.POST("/users/new/:id", posts)
    e.DELETE("/users/delete/:id", deletes)
	//logger and server start
	e.Logger.Fatal(e.Start(":4000"))

}


