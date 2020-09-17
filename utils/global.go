package utils

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io/ioutil"
)

type APIResponse struct {
	Status    string      `json:"status"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	ErrorCode string      `json:"errorCode,omitempty"`
	Total     int64       `json:"total,omitempty"`
}

type StatusEnum struct {
	Ok           string
	Error        string
	Invalid      string
	NotFound     string
	Forbidden    string
	Existed      string
	Unauthorized string
}

var APIStatus = &StatusEnum{
	Ok:           "OK",
	Error:        "ERROR",
	Invalid:      "INVALID",
	NotFound:     "NOT_FOUND",
	Forbidden:    "FORBIDDEN",
	Existed:      "EXISTED",
	Unauthorized: "UNAUTHORIZED",
}

type MgoOperation struct {
	Or    interface{} `bson:"$or,omitempty"`
	In    interface{} `bson:"$in,omitempty"`
	Gt    interface{} `bson:"$gt,omitempty"`
	Gte   interface{} `bson:"$gte,omitempty"`
	Eq    interface{} `bson:"$eq,omitempty"`
	Ne    interface{} `bson:"$ne,omitempty"`
	Lt    interface{} `bson:"$lt,omitempty"`
	Lte   interface{} `bson:"$lte,omitempty"`
	Match interface{} `bson:"$match,omitempty"`
	Group interface{} `bson:"$group,omitempty"`
	Sum   interface{} `bson:"$sum,omitempty"`
	Text  string      `bson:"$search,omitempty"`
	Regex string      `bson:"$regex,omitempty"`
}

func GetContent(c echo.Context, input interface{}) error {
	body := c.Request().Body
	var bodyBytes []byte
	if body != nil {
		bodyBytes, _ = ioutil.ReadAll(body)
	}
	b := string(bodyBytes)
	return json.Unmarshal([]byte(b), &input)
}
