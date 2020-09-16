package model

import (
	"github.com/2D03/comtam-be/utils"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"time"
)

type Order struct {
	ID              bson.ObjectId `json:"-" bson:"_id,omitempty"`
	CreatedTime     *time.Time    `json:"createdTime" bson:"created_time,omitempty"`
	LastUpdatedTime *time.Time    `json:"lastUpdatedTime" bson:"last_updated_time,omitempty"`
	TotalPrice      *int64        `json:"totalPrice" bson:"total_price,omitempty"`
	Dishes          []*OrderDish  `json:"dishes" bson:"dishes,omitempty"`
	Address         *string       `json:"address" bson:"address,omitempty"`
	Phone           *string       `json:"phone" bson:"phone,omitempty"`
	OrdererName     *string       `json:"ordererName" bson:"orderer_name,omitempty"`
	ReceiverName    *string       `json:"receiverName" bson:"receiver_name,omitempty"`
}

type OrderDish struct {
	UniqueID *string `json:"uniqueId,omitempty" bson:"unique_id,omitempty"`
	Name     *string `json:"name" bson:"name,omitempty"`
	MenuId   *string `json:"menuId" bson:"menu_id,omitempty"`
	Price    *int64  `json:"price" bson:"price,omitempty"`
	Amount   *int    `json:"amount" bson:"amount,omitempty"`
}

var OrderModel = &utils.DBModel{
	ColName:       "orders",
	TemplateModel: &Order{},
}

func InitOrderModel(s *mgo.Session, dbName string) {
	OrderModel.DBName = dbName
	err := OrderModel.Init(s)
	if err != nil {
		panic(err)
	}
	_ = OrderModel.CreateIndex(mgo.Index{
		Key:        []string{"total_price"},
		Background: true, // See notes.
	})
}
