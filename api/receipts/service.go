package receipts

import (
	"net/http"

	"github.com/anisbhsl/fetch-backend-assignment/store"
)

type Service interface {
	ProcessReceipts() http.HandlerFunc
	ProcessReceiptPoints() http.HandlerFunc
}

// service implements receipt Service interface
type service struct {
	db store.Service
}

func NewReceiptsAPIService(db store.Service) Service {
	return &service{
		db: db,
	}
}
