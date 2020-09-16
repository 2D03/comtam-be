package api

import (
	"github.com/2D03/comtam-be/model"
	"github.com/2D03/comtam-be/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

func UpdateMenu(c echo.Context) error {
	var input *model.ReqUpdateMenu
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
	query := model.Menu{
		UniqueID: input.UniqueID,
	}
	var updater model.Menu
	if input.Name != nil {
		updater.Name = input.Name
	}
	var queryDishesResult *utils.APIResponse
	if len(input.Dishes) > 0 {
		queryDishesResult = model.DishModel.Query(model.DishForFilter{
			UniqueID: utils.MgoOperation{In: input.Dishes},
		})
		if queryDishesResult.Status != utils.APIStatus.Ok {
			return c.JSON(http.StatusInternalServerError, &utils.APIResponse{
				Status:  queryDishesResult.Status,
				Message: "Cannot get dishes info: " + queryDishesResult.Message,
			})
		}
		updater.Dish = queryDishesResult.Data.([]*model.Dish)
	}

	res := model.MenuModel.Update(query, updater)
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
