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

	t, err := bson.Marshal(obj)
	if err != nil {
		return &APIResponse{
			Status:  APIStatus.Error,
			Message: "Marshal error: " + err.Error(),
		}
	}
	typ := reflect.TypeOf(m.TemplateModel)
	createdObj := reflect.New(typ).Interface()
	err = bson.Unmarshal(t, createdObj)
	if err != nil {
		return &APIResponse{
			Status:  APIStatus.Error,
			Message: "Unmarshal error: " + err.Error(),
		}
	}

	return &APIResponse{
		Status:  APIStatus.Ok,
		Message: "Create " + m.ColName + " successfully.",
		Data:    createdObj,
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
	if err != nil || reflect.ValueOf(list).Len() == 0 {
		return &APIResponse{
			Status:  APIStatus.NotFound,
			Message: "Not found any matched " + m.ColName + ".",
		}
	}
	return &APIResponse{
		Status:  APIStatus.Ok,
		Message: "Query " + m.ColName + " successfully.",
		Data:    list,
	}
}

func (m *DBModel) Update(query interface{}, updater interface{}) *APIResponse {
	s := m.session.Copy()
	defer s.Close()
	if m.collection == nil {
		m.collection = s.DB(m.DBName).C(m.ColName)
	}
	col := m.collection.With(s)
	obj, err := m.convertToBson(query)
	if err != nil {
		return &APIResponse{
			Status:  APIStatus.Error,
			Message: "DB error: " + err.Error(),
		}
	}
	obj["last_updated_time"] = time.Now()
	info, err := col.UpdateAll(query, bson.M{
		"$set": obj,
	})
	if err != nil {
		return &APIResponse{
			Status:    APIStatus.Error,
			Message:   "Update error: " + err.Error(),
			ErrorCode: "UPDATE_FAILED",
		}
	}
	if info.Matched == 0 {
		return &APIResponse{
			Status:  APIStatus.Ok,
			Message: "Not found any " + m.ColName + ".",
		}
	}

	return &APIResponse{
		Status:  APIStatus.Ok,
		Message: "Update " + m.ColName + " successfully.",
	}
}

func (m *DBModel) Delete(selector interface{}) *APIResponse {
	s := m.session.Copy()
	defer s.Close()
	if m.collection == nil {
		m.collection = s.DB(m.DBName).C(m.ColName)
	}
	col := m.collection.With(s)
	err := col.Remove(selector)
	if err != nil {
		return &APIResponse{
			Status:    APIStatus.Error,
			Message:   "Delete error: " + err.Error(),
			ErrorCode: "DELETE_FAILED",
		}
	}
	return &APIResponse{
		Status:  APIStatus.Ok,
		Message: "Delete " + m.ColName + " successfully.",
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
