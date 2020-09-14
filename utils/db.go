package utils

import (
	"errors"
	"github.com/globalsign/mgo"
)

type DBModel struct {
	ColName    string
	DBName     string
	collection *mgo.Collection
	session    *mgo.Session
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
