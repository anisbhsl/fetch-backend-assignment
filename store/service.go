package store

import "github.com/anisbhsl/fetch-backend-assignment/models"

type Service interface {
	Put(receipt models.Receipt) string
	Get() models.Receipt
}
