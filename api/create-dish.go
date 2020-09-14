package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateDish(c echo.Context) error {
	return c.String(http.StatusOK, "Service run normally")
}
