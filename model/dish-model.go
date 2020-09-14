package model

import (
	"github.com/Kajekk/comtam/utils"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"time"
)

type Dish struct {
	ID              bson.ObjectId `json:"-" bson:"_id,omitempty"`
	CreatedTime     *time.Time    `json:"createdTime" bson:"created_time,omitempty"`
	LastUpdatedTime *time.Time    `json:"lastUpdatedTime" bson:"last_updated_time,omitempty"`
	CreatedById     *string       `json:"createdById" bson:"created_by_id,omitempty"`
	CreatedByName   *string       `json:"createdByName" bson:"created_by_name,omitempty"`
	UniqueID        *string       `json:"uniqueId,omitempty" bson:"unique_id,omitempty"`
	PriceAmount     *int64        `json:"priceAmount" bson:"price_amount,omitempty"`
}

var DishModel = &utils.DBModel{
	ColName: "dishes",
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
