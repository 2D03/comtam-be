package api

import (
	"github.com/2D03/comtam-be/model"
	"github.com/2D03/comtam-be/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

func UpdateDish(c echo.Context) error {
	var input *model.Dish
	err := utils.GetContent(c, &input)
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
	query := model.Dish{
		UniqueID: input.UniqueID,
	}
	var updater model.Dish
	if input.Name != nil {
		updater.Name = input.Name
	}
	if input.PriceAmount != nil {
		updater.PriceAmount = input.PriceAmount
	}
	if input.MenuId != nil {
		updater.MenuId = input.MenuId
	}
	res := model.DishModel.Update(query, updater)
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
