package api

import (
	"github.com/2D03/comtam-be/model"
	"github.com/2D03/comtam-be/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetMenu(c echo.Context) error {
	id := c.QueryParam("uniqueId")
	var query model.Menu
	if id != "" {
		query.UniqueID = &id
	}
	rs := model.MenuModel.Query(query)
	if rs.Status != utils.APIStatus.Ok {
		return c.JSON(http.StatusInternalServerError, &utils.APIResponse{
			Status:  utils.APIStatus.Error,
			Message: rs.Message,
		})
	}
	return c.JSON(http.StatusOK, &utils.APIResponse{
		Status:  rs.Status,
		Message: rs.Message,
		Data:    rs.Data,
		Total:   rs.Total,
	})
}
