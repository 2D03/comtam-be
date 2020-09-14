package api

import (
	"encoding/json"
	"github.com/Kajekk/comtam/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetMenu(c echo.Context) error {
	q := c.QueryParam("q")
	var input *model.ReqGetMenu
	err := json.Unmarshal([]byte(q), &input)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error")
	}
	return c.String(http.StatusOK, "Hello, World!")
}
