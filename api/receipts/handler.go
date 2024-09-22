package receipts

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/anisbhsl/fetch-backend-assignment/models"
	"github.com/anisbhsl/fetch-backend-assignment/utils"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

// ProcessReceipts takes in a JSON payload for a receipt, saves it to db
// and return an ID
func (s service) ProcessReceipts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.GetLogger().Info("processing receipts")

		var receiptRequest models.Receipt
		if err := json.NewDecoder(r.Body).Decode(&receiptRequest); err != nil {
			utils.GetLogger().Error("error while decoding receipt")
			utils.SendErrorResponse(w, "Invalid Receipt", http.StatusBadRequest)
			return
		}

		utils.GetLogger().Debug(fmt.Sprintf("incoming payload is: %v", receiptRequest))

		//validate the payload
		if err := utils.GetValidator().Struct(receiptRequest); err != nil {
			utils.GetLogger().Error("payload invalid")
			utils.SendErrorResponse(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
			return
		}

		// save the receipt to storage
		receiptID, _ := s.db.StoreReceipt(&receiptRequest)
		utils.GetLogger().Info("stored receipt into db", zap.String("id", receiptID))

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
		vars := mux.Vars(r)
		receiptID := vars["id"]

		utils.GetLogger().Info("calculating points for given receipt id", zap.String("id", receiptID))

		// validate receipt ID
		receiptIDValidator := regexp.MustCompile(`^\S+$`)
		if !receiptIDValidator.MatchString(receiptID) {
			utils.GetLogger().Error("invalid receipt id format")
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
			utils.GetLogger().Error("error while calculating points", zap.Error(err))
			utils.SendErrorResponse(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
			return
		}

		utils.GetLogger().Info(fmt.Sprintf("the total points for the receipt is: %v", totalPoints))

		// send response back to client
		utils.SendSuccessResponse(w, models.CalculatePointsResponse{
			Points: totalPoints,
		})
	}
}
