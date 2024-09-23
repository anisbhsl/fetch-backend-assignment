package receipts

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/anisbhsl/fetch-backend-assignment/mock"
	"github.com/anisbhsl/fetch-backend-assignment/models"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

var PROCESS_RECEIPTS_ENDPOINT = "/api/v1/receipts/process"

// TestProcessReceiptsNilItems tests by removing items from the payload
// to check if validator catches it or not.
func TestProcessReceiptsNilItems(t *testing.T) {
	receipt := &models.Receipt{
		ID:           "ca30cf14-952c-4e0b-8dbe-13e0f602b87a",
		Retailer:     "Walmart",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Total:        "35.00",
	}

	store := mock.Store{
		Receipt: receipt,
	}

	service := NewReceiptsAPIService(store)
	r := mux.NewRouter()
	r.HandleFunc(PROCESS_RECEIPTS_ENDPOINT, service.ProcessReceipts()).Methods("POST")

	bodyInBytes, err := json.Marshal(receipt)
	if err != nil {
		t.Fatalf("failed to marshal receipt: %v", err)
	}

	in := httptest.NewRequest("POST", PROCESS_RECEIPTS_ENDPOINT, strings.NewReader(string(bodyInBytes)))
	in.Header.Set("Content-Type", "application/json")

	out := httptest.NewRecorder()

	r.ServeHTTP(out, in)
	assert.Equal(t, 400, out.Code)
}

// TestProcessReceiptsForTarget tests by passing the target receipt
// this test should send 200 and the id as response
func TestProcessReceiptsForTarget(t *testing.T) {
	receipt := &models.Receipt{
		ID:           "ca30cf14-952c-4e0b-8dbe-13e0f602b87a",
		Retailer:     "Target",
		PurchaseDate: "2022-01-02",
		PurchaseTime: "13:13",
		Total:        "1.25",
		Items: []*models.ReceiptItems{
			{
				ShortDescription: "Pepsi - 12-oz",
				Price:            "1.25",
			},
		},
	}

	store := mock.Store{
		Receipt: receipt,
	}

	service := NewReceiptsAPIService(store)
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/receipts", service.ProcessReceipts()).Methods("POST")
	out := httptest.NewRecorder()
	bodyInBytes, _ := json.Marshal(receipt)

	in := httptest.NewRequest("POST", PROCESS_RECEIPTS_ENDPOINT, strings.NewReader(string(bodyInBytes)))
	r.ServeHTTP(out, in)
	assert.Equal(t, 200, out.Code)
	assert.JSONEq(t, `{"id": "ca30cf14-952c-4e0b-8dbe-13e0f602b87a"}`, out.Body.String())
}

// TestProcessReceiptPointsForWalgreens calculates the points for walgreens receipt
func TestProcessReceiptPointsForWalgreens(t *testing.T) {
	receiptID := "ca30cf14-952c-4e0b-8dbe-13e0f602b87a"
	receipt := &models.Receipt{
		ID:           receiptID,
		Retailer:     "Walgreens",
		PurchaseDate: "2022-01-02",
		PurchaseTime: "08:13",
		Total:        "2.65",
		Items: []*models.ReceiptItems{
			{
				ShortDescription: "Pepsi - 12-oz",
				Price:            "1.25",
			},
			{
				ShortDescription: "Dasani",
				Price:            "1.40",
			},
		},
	}

	store := mock.Store{
		Receipt: receipt,
	}

	service := NewReceiptsAPIService(store)
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/receipts/{id}/points", service.ProcessReceiptPoints()).Methods("GET")
	out := httptest.NewRecorder()
	in := httptest.NewRequest("GET", fmt.Sprintf("/api/v1/receipts/%s/points", receiptID), nil)
	r.ServeHTTP(out, in)
	assert.Equal(t, 200, out.Code)

	// t.Log(out.Body.String())
	assert.JSONEq(t, `{"points":15}`, out.Body.String())
}
