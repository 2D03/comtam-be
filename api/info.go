package api

import (
	"github.com/2D03/comtam-be/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetAPIInfo(c echo.Context) error {
	return c.JSON(http.StatusOK, &utils.APIResponse{
		Status:  utils.APIStatus.Ok,
		Message: "Service run successfully",
	})
}
