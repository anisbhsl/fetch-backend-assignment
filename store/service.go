package store

import "github.com/anisbhsl/fetch-backend-assignment/models"

type Service interface {
	StoreReceipt(receipt *models.Receipt) (string, error)
	GetReceipt(id string) (*models.Receipt, error)
	StoreReceiptPoints(id string, total_points int) error
	GetReceiptPoints(id string) (int, error)
}
