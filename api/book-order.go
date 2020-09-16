package api

import (
	"github.com/2D03/comtam-be/conf"
	"github.com/2D03/comtam-be/model"
	"github.com/2D03/comtam-be/utils"
	"github.com/labstack/echo/v4"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"net/http"
	"strconv"
)

func BookOrder(c echo.Context) error {
	var input *model.Order
	err := utils.GetContent(c, &input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &utils.APIResponse{
			Status:  utils.APIStatus.Error,
			Message: "Error: " + err.Error(),
		})
	}

	if input.Address == nil || *input.Address == "" {
		return c.JSON(http.StatusBadRequest, &utils.APIResponse{
			Status:    utils.APIStatus.Invalid,
			Message:   "Missing Address",
			ErrorCode: "MISSING_ADDRESS",
		})
	}
	if input.Phone == nil || *input.Phone == "" {
		return c.JSON(http.StatusBadRequest, &utils.APIResponse{
			Status:    utils.APIStatus.Invalid,
			Message:   "Missing Phone",
			ErrorCode: "MISSING_PHONE",
		})
	}
	if input.ReceiverName == nil || *input.ReceiverName == "" {
		return c.JSON(http.StatusBadRequest, &utils.APIResponse{
			Status:    utils.APIStatus.Invalid,
			Message:   "Missing Name",
			ErrorCode: "MISSING_NAME",
		})
	}
	if input.Dishes == nil || len(input.Dishes) == 0 {
		return c.JSON(http.StatusBadRequest, &utils.APIResponse{
			Status:    utils.APIStatus.Invalid,
			Message:   "Missing Dishes",
			ErrorCode: "MISSING_DISHES",
		})
	}
	defaultPrice := int64(0)
	if input.TotalPrice == nil {
		input.TotalPrice = &defaultPrice
	}

	rs := sendEmail(input)
	if rs.Status != utils.APIStatus.Ok {
		return c.JSON(http.StatusInternalServerError, &utils.APIResponse{
			Status:  rs.Status,
			Message: rs.Message,
		})
	}

	createResult := model.OrderModel.Create(input)
	if createResult.Status != utils.APIStatus.Ok {
		return c.JSON(http.StatusOK, &utils.APIResponse{
			Status:  rs.Status,
			Message: "Error when creating OrderDB: " + rs.Message,
		})
	}

	return c.JSON(http.StatusOK, &utils.APIResponse{
		Status:  utils.APIStatus.Ok,
		Message: "Booking order successfully",
	})
}

func sendEmail(input *model.Order) *utils.APIResponse {
	request := sendgrid.GetRequest(conf.SendgridAPIKey, "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	var Body = infoMail(input)
	request.Body = Body
	response, err := sendgrid.API(request)
	if err != nil {
		return &utils.APIResponse{
			Status:  utils.APIStatus.Error,
			Message: "Error: " + err.Error(),
		}
	} else {
		return &utils.APIResponse{
			Status:  utils.APIStatus.Ok,
			Message: response.Body,
		}
	}
}

func infoMail(input *model.Order) []byte {
	address := "test@example.com"
	name := "Comtam-Mail-Sender"
	from := mail.NewEmail(name, address)
	subject := "Khách hàng vừa đặt món"
	dishes := ""
	for _, dish := range input.Dishes {
		if dish.Amount != nil && dish.Name != nil && *dish.Amount > 0 && *dish.Name != "" {
			dishes = dishes + *dish.Name + " - Số lượng: " + strconv.Itoa(*dish.Amount) + "\n"
		}
	}
	plainTextContent := "Thông tin đơn hàng:\nTên người nhận: " + *input.ReceiverName + "\n" +
		"Số điện thoại người nhận: " + *input.Phone + "\n" +
		"Địa chỉ: " + *input.Address + "\n" +
		"Tổng tiền: " + strconv.FormatInt(*input.TotalPrice, 10) + "\n" +
		"Đơn hàng: \n" + dishes
	content := mail.NewContent("text/plain", plainTextContent)

	address = conf.ToEmail[0]
	name = conf.ToEmail[0]
	to := mail.NewEmail(name, address)
	m := mail.NewV3MailInit(from, subject, to, content)
	for i, toMail := range conf.ToEmail {
		if toMail != "" && i > 0 {
			email := mail.NewEmail(toMail, toMail)
			m.Personalizations[0].AddTos(email)
		}
	}

	return mail.GetRequestBody(m)
}
