package api

import (
	"encoding/json"
	"github.com/2D03/comtam-be/model"
	"github.com/2D03/comtam-be/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateMenu(c echo.Context) error {
	q := c.QueryParam("q")
	var input model.Menu
	err := json.Unmarshal([]byte(q), &input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &utils.APIResponse{
			Status:  utils.APIStatus.Error,
			Message: "Error: " + err.Error(),
		})
	}

	if input.UniqueID == nil {
		return c.JSON(http.StatusBadRequest, &utils.APIResponse{
			Status:  utils.APIStatus.Invalid,
			Message: "Missing UniqueId",
		})
	}

	result := model.MenuModel.Create(input)
	if result.Status != utils.APIStatus.Ok {
		return c.JSON(http.StatusInternalServerError, &utils.APIResponse{
			Status:  utils.APIStatus.Error,
			Message: result.Message,
		})
	}

	return c.JSON(http.StatusOK, &utils.APIResponse{
		Status:  result.Status,
		Message: result.Message,
		Data:    result.Data,
	})
}
