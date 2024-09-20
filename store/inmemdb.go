package store

import "github.com/anisbhsl/fetch-backend-assignment/models"

type inMemDB struct {
}

func NewInMemDB() Service {
	s := &inMemDB{}
	return s
}

func (db inMemDB) Put(receipt models.Receipt) string {
	return ""
}

func (db inMemDB) Get() models.Receipt {
	return models.Receipt{}
}
