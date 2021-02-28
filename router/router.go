package router

import (
	"api/models"
	"database/sql"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type H map[string]interface{}

func GetUser(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		// Fetch comments using our new model
		return c.JSON(http.StatusOK, models.GetUsers(db, id))
	}
}

func PutUser(c echo.Context) error {
	id := c.Param("id")
	sum := id + "Put user"
	return c.String(http.StatusOK, sum)
}
/* bug TO do
func PostUser(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
        db := models.InitDB("../sqlite.db")
		models.AddUser(db, id, "new")

		return c.String(http.StatusOK,"OKEY")
	}
}
*/


func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
