package api

import (
	"encoding/json"
	"github.com/2D03/comtam-be/model"
	"github.com/2D03/comtam-be/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateDish(c echo.Context) error {
	q := c.QueryParam("q")
	var input model.Dish
	_ = json.Unmarshal([]byte(q), &input)

	result := model.DishModel.Create(input)
	if result.Status != utils.APIStatus.Ok {
		return c.String(http.StatusInternalServerError, result.Message)
	}

	//return &utils.APIResponse{
	//	Status:  utils.APIStatus.Ok,
	//	Message: "Created dish successfully",
	//}

	return c.String(http.StatusOK, result.Message)
}
