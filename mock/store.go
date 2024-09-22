/*
This package serves as a mock for store service
*/
package mock

import "github.com/anisbhsl/fetch-backend-assignment/models"

// Store is a mock implementation of store
// for testing purposes
type Store struct {
	Receipt *models.Receipt
}

func (s Store) StoreReceipt(receipt *models.Receipt) (string, error) {
	return s.Receipt.ID, nil

}
func (s Store) GetReceipt(id string) (*models.Receipt, error) {
	return s.Receipt, nil
}

func (s Store) StoreReceiptPoints(id string, total_points int) error {
	return nil
}

func (s Store) GetReceiptPoints(id string) (int, error) {
	return 0, nil
}
