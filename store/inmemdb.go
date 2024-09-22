package store

import (
	"fmt"

	"github.com/anisbhsl/fetch-backend-assignment/models"
	"github.com/google/uuid"
)

// inMemDB implements a simple in-memory key value store
/*
Some key naming patterns followed:

1. Receipts are saved as=> receipt::receipt_id
2. Store receipt points as=> receipt::receipt_id::points
*/
type inMemDB struct {
	kvStore map[string]interface{}
}

func NewInMemDB() Service {
	s := &inMemDB{
		kvStore: make(map[string]interface{}),
	}
	return s
}

// StoreReceipt saves the receipt to database and returns UUID based unique ID
// for each receipt
func (db inMemDB) StoreReceipt(receipt *models.Receipt) (string, error) {
	id := uuid.New().String()

	// save the receipt
	receipt.ID = id

	key := fmt.Sprintf("receipt::%s", id)
	db.kvStore[key] = receipt
	return id, nil
}

// GetReceipt retrieves the receipt for a given ID if it exists other will return nil
func (db inMemDB) GetReceipt(id string) (*models.Receipt, error) {
	key := fmt.Sprintf("receipt::%s", id)
	receipt, ok := db.kvStore[key]
	if !ok {
		return nil, nil
	}

	return receipt.(*models.Receipt), nil
}

// StoreReceiptPoints stores the points to a given receipt
func (db inMemDB) StoreReceiptPoints(id string, total_points int) error {
	key := fmt.Sprintf("receipt::%s::points", id)
	db.kvStore[key] = total_points
	return nil
}

// GetReceiptPoints retrievs the total points awarded to a given receipt
func (db inMemDB) GetReceiptPoints(id string) (int, error) {
	key := fmt.Sprintf("receipt::%s::points", id)
	total_points, ok := db.kvStore[key]
	if !ok {
		return -1, nil
	}
	return total_points.(int), nil
}
