package model

import (
	"github.com/2D03/comtam-be/utils"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"time"
)

type Menu struct {
	ID              bson.ObjectId `json:"-" bson:"_id,omitempty"`
	CreatedTime     *time.Time    `json:"createdTime" bson:"created_time,omitempty"`
	LastUpdatedTime *time.Time    `json:"lastUpdatedTime" bson:"last_updated_time,omitempty"`
	UniqueID        *string       `json:"uniqueId,omitempty" bson:"unique_id,omitempty"`
	Name            *string       `json:"name,omitempty" bson:"name,omitempty"`
}

type ReqUpdateMenu struct {
	UniqueID *string `json:"uniqueId,omitempty"`
	Name     *string `json:"name,omitempty"`
}

type MenuForFilter struct {
	CreatedById   interface{} `json:"createdById,omitempty" bson:"created_by_id,omitempty"`
	CreatedByName interface{} `json:"createdByName,omitempty" bson:"created_by_name,omitempty"`
	UniqueID      interface{} `json:"uniqueId,omitempty" bson:"unique_id,omitempty"`
	Name          interface{} `json:"name,omitempty" bson:"name,omitempty"`
}

var MenuModel = &utils.DBModel{
	ColName:       "menu",
	TemplateModel: Menu{},
}

func InitMenuModel(s *mgo.Session, dbName string) {
	MenuModel.DBName = dbName
	err := MenuModel.Init(s)
	if err != nil {
		panic(err)
	}
	_ = MenuModel.CreateIndex(mgo.Index{
		Key:        []string{"unique_id"},
		Unique:     true,
		Background: true, // See notes.
	})
}
