package api

import (
	"github.com/2D03/comtam-be/model"
	"github.com/2D03/comtam-be/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

func DeleteMenu(c echo.Context) error {
	uniqueId := c.QueryParam("uniqueId")

	if uniqueId == "" {
		return c.JSON(http.StatusBadRequest, &utils.APIResponse{
			Status:  utils.APIStatus.Invalid,
			Message: "Missing UniqueId",
		})
	}
	selector := model.Menu{
		UniqueID: &uniqueId,
	}

	res := model.MenuModel.Delete(selector)
	if res.Status != utils.APIStatus.Ok {
		return c.JSON(http.StatusInternalServerError, &utils.APIResponse{
			Status:  res.Status,
			Message: res.Message,
		})
	}
	return c.JSON(http.StatusOK, &utils.APIResponse{
		Status:  res.Status,
		Message: res.Message,
	})
}
