package model

import (
	"github.com/2D03/comtam-be/utils"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"time"
)

type Dish struct {
	ID              bson.ObjectId `json:"-" bson:"_id,omitempty"`
	CreatedTime     *time.Time    `json:"createdTime" bson:"created_time,omitempty"`
	LastUpdatedTime *time.Time    `json:"lastUpdatedTime" bson:"last_updated_time,omitempty"`
	UniqueID        *string       `json:"uniqueId,omitempty" bson:"unique_id,omitempty"`
	Name            *string       `json:"name" bson:"name,omitempty"`
	MenuId          *string       `json:"menuId" bson:"menu_id,omitempty"`
	PriceAmount     *int64        `json:"priceAmount" bson:"price_amount,omitempty"`
}

type DishForFilter struct {
	CreatedById   interface{} `json:"createdById" bson:"created_by_id,omitempty"`
	CreatedByName interface{} `json:"createdByName" bson:"created_by_name,omitempty"`
	UniqueID      interface{} `json:"uniqueId" bson:"unique_id,omitempty"`
	MenuId        interface{} `json:"menuId" bson:"menu_id,omitempty"`
	Name          interface{} `json:"name" bson:"name,omitempty"`
	PriceAmount   interface{} `json:"priceAmount" bson:"price_amount,omitempty"`
}

type ReqDish struct {
	UniqueID    *string `json:"uniqueId"`
	Name        *string `json:"name"`
	PriceAmount *int64  `json:"priceAmount"`
}

var DishModel = &utils.DBModel{
	ColName:       "dishes",
	TemplateModel: &Dish{},
}

func InitDishModel(s *mgo.Session, dbName string) {
	DishModel.DBName = dbName
	err := DishModel.Init(s)
	if err != nil {
		panic(err)
	}
	_ = DishModel.CreateIndex(mgo.Index{
		Key:        []string{"unique_id"},
		Unique:     true,
		Background: true, // See notes.
	})
}
