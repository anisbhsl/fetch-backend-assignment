package models

type ReceiptItems struct {
	ShortDescription string `json:"shortDescription" validate:"required"`
	Price            string `json:"Price" validate:"required,numeric"`
}

type Receipt struct {
	Retailer     string          `json:"retailer" validate:"required"`
	PurchaseDate string          `json:"purchaseDate" validate:"required,datetime=2006-01-02"`
	PurchaseTime string          `json:"purchaseTime" validate:"required,datetime=15:04"`
	Items        []*ReceiptItems `json:"items" validate:"required,dive"`
	Total        string          `json:"total" validate:"required,numeric,gte=0"`
}

type ProcessReceiptResponse struct {
	ID string `json:"id"`
}

type CalculatePointsResponse struct {
	Points string `json:"points"`
}
