package receipts

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/anisbhsl/fetch-backend-assignment/models"
	"github.com/anisbhsl/fetch-backend-assignment/utils"
	"github.com/gorilla/mux"
)

// ProcessReceipts takes in a JSON payload for a receipt, saves it to db
// and return an ID
func (s service) ProcessReceipts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		print("processing receipts....")
		var receiptRequest models.Receipt
		if err := json.NewDecoder(r.Body).Decode(&receiptRequest); err != nil {
			// TODO: throw error log
			utils.SendErrorResponse(w, "Invalid Receipt", http.StatusBadRequest)
			return
		}

		//validate the payload
		if err := utils.GetValidator().Struct(receiptRequest); err != nil {
			//TODO: throw error log
			utils.SendErrorResponse(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
			return
		}

		// save the receipt to storage
		receiptID, _ := s.db.StoreReceipt(&receiptRequest)

		//send back the response
		utils.SendSuccessResponse(w, models.ProcessReceiptResponse{
			ID: receiptID,
		})
	}
}

// ProcessReceiptPoints takes receipt ID to calculate points awarded for a
// given receipt
func (s service) ProcessReceiptPoints() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		print("handling receipt points")

		vars := mux.Vars(r)
		receiptID := vars["id"]

		// validate receipt ID
		receiptIDValidator := regexp.MustCompile(`^\S+$`)
		if !receiptIDValidator.MatchString(receiptID) {
			utils.SendErrorResponse(w, "Invalid receipt ID format", http.StatusBadRequest)
			return
		}

		//retrieve receipt from the db if it exists
		receipt, err := s.db.GetReceipt(receiptID)
		if err != nil {
			utils.SendErrorResponse(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
			return
		} else if receipt == nil {
			utils.SendErrorResponse(w, "No receipt found for that id", http.StatusNotFound)
			return
		}

		//calculate points for given receipt
		totalPoints, err := receipt.CalculatePoints()
		if err != nil {
			utils.SendErrorResponse(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
			return
		}

		// send response back to client
		utils.SendSuccessResponse(w, models.CalculatePointsResponse{
			Points: totalPoints,
		})
	}
}
