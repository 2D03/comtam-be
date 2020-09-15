package utils

import (
	"errors"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"reflect"
	"time"
)

type DBModel struct {
	TemplateModel interface{}
	ColName       string
	DBName        string
	collection    *mgo.Collection
	session       *mgo.Session
}

func (m *DBModel) Init(s *mgo.Session) error {
	if len(m.DBName) == 0 || len(m.ColName) == 0 {
		return errors.New("Require valid DB name and collection name.")
	}

	m.session = s
	m.collection = s.DB(m.DBName).C(m.ColName)
	return nil
}

func (m *DBModel) CreateIndex(index mgo.Index) error {
	s := m.session.Copy()
	defer s.Close()
	if m.collection == nil {
		m.collection = s.DB(m.DBName).C(m.ColName)
	}
	col := m.collection.With(s)
	err := col.EnsureIndex(index)
	if err != nil {
		return err
	}
	return nil
}

func (m *DBModel) Create(entity interface{}) *APIResponse {
	s := m.session.Copy()
	defer s.Close()
	if m.collection == nil {
		m.collection = s.DB(m.DBName).C(m.ColName)
	}
	col := m.collection.With(s)
	//err := col.Insert(index)
	obj, err := m.convertToBson(entity)
	if err != nil {
		return &APIResponse{
			Status:  APIStatus.Error,
			Message: "DB error: " + err.Error(),
		}
	}

	if obj["created_time"] == nil {
		obj["created_time"] = time.Now()
		obj["last_updated_time"] = obj["created_time"]
	} else {
		obj["last_updated_time"] = time.Now()
	}

	err = col.Insert(obj)
	if err != nil {
		return &APIResponse{
			Status:  APIStatus.Error,
			Message: "DB error: " + err.Error(),
		}
	}

	return &APIResponse{
		Status:  APIStatus.Ok,
		Message: "Create " + m.ColName + " successfully.",
	}
}

func (m *DBModel) Query(query interface{}) *APIResponse {
	s := m.session.Copy()
	defer s.Close()
	if m.collection == nil {
		m.collection = s.DB(m.DBName).C(m.ColName)
	}
	col := m.collection.With(s)
	q := col.Find(query)
	t := reflect.TypeOf(m.TemplateModel)
	list := reflect.MakeSlice(reflect.SliceOf(t), 0, 0).Interface()
	err := q.All(&list)
	if err != nil {
		return &APIResponse{
			Status:  APIStatus.Error,
			Message: "Error: " + err.Error(),
		}
	}
	return &APIResponse{
		Status:  APIStatus.Ok,
		Message: "Query " + m.ColName + " successfully.",
		Data:    list,
	}
}

func (m *DBModel) convertToBson(entity interface{}) (bson.M, error) {
	if entity == nil {
		return bson.M{}, nil
	}

	temp, err := bson.Marshal(entity)
	if err != nil {
		return nil, err
	}

	obj := bson.M{}
	err = bson.Unmarshal(temp, &obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
