package receipts

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/anisbhsl/fetch-backend-assignment/models"
	"github.com/anisbhsl/fetch-backend-assignment/utils"
	// "github.com/go-playground/validator/v10"
)

func (s service) ProcessReceipts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		print("processing receipts....")
		var receiptRequest models.Receipt
		err := json.NewDecoder(r.Body).Decode(&receiptRequest)
		if err != nil {
			// TODO: throw error log
			utils.SendErrorResponse(w, "Invalid Receipt")
			return
		}

		//validate the payload
		if err := utils.Validate.Struct(receiptRequest); err != nil {
			//TODO: throw error log
			utils.SendErrorResponse(w, fmt.Sprintf("%v", err))
			return
		}

		// TODO: add logic

		utils.SendSuccessResponse(w, models.ProcessReceiptResponse{
			ID: "123",
		})
	}
}

func (s service) ProcessReceiptPoints() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// var receipt models.Receipt
		// err := json.NewDecoder(r.Body).Decode(&receipt)
		// if err != nil {
		// 	// TODO: throw error log
		// 	utils.SendErrorResponse(w, "Invalid Receipt")
		// 	return
		// }
		print("handling receipt points")

		utils.SendSuccessResponse(w, models.ProcessReceiptResponse{
			ID: "123",
		})
	}
}
