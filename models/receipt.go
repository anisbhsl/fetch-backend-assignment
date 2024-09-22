package models

import (
	"math"
	"regexp"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
)

type ReceiptItems struct {
	ShortDescription string `json:"shortDescription" validate:"required,validateReceiptItemShortDesc"`
	Price            string `json:"price" validate:"required,validateReceiptItemPrice"`
}

type Receipt struct {
	ID           string          `json:"-"`
	Retailer     string          `json:"retailer" validate:"required,validateRetailer"`
	PurchaseDate string          `json:"purchaseDate" validate:"required,datetime=2006-01-02"`
	PurchaseTime string          `json:"purchaseTime" validate:"required,datetime=15:04"`
	Items        []*ReceiptItems `json:"items" validate:"required,min=1,dive,required"`
	Total        string          `json:"total" validate:"required,validateReceiptTotal"`
}

func validateAgainstRegex(pattern string, fieldToValidate string) bool {
	matched, _ := regexp.MatchString(pattern, fieldToValidate)
	return matched
}

// ValidateRetailer is a custom field validator for retailer name
// uses ^[\\w\\s\\-&]+$ pattern to validate
func ValidateRetailer(f validator.FieldLevel) bool {
	retailer := f.Field().String()
	pattern := `^[\w\s\-&]+$`
	return validateAgainstRegex(pattern, retailer)
}

// ValidateReceiptTotal validates the total of a receipts against
// a regex provided.
func ValidateReceiptTotal(f validator.FieldLevel) bool {
	total := f.Field().String()
	pattern := `^\d+\.\d{2}$`
	return validateAgainstRegex(pattern, total)
}

// ValidateReceiptItemShortDesc validates short description field
// for receipt item against a regex
func ValidateReceiptItemShortDesc(f validator.FieldLevel) bool {
	desc := f.Field().String()
	pattern := `^[\w\s\-]+$`
	return validateAgainstRegex(pattern, desc)
}

// ValidateReceiptItemPrice validates the receipt item price against a regex
func ValidateReceiptItemPrice(f validator.FieldLevel) bool {
	price := f.Field().String()
	pattern := `\d+\.\d{2}$`
	return validateAgainstRegex(pattern, price)
}

// isPurchasedDateOdd checks if the given day is odd or not
func (receipt *Receipt) isPurchasedDateOdd() (bool, error) {
	dateFormat := "2006-01-02"
	parsedDate, err := time.Parse(dateFormat, receipt.PurchaseDate)
	if err != nil {
		return false, err
	}
	purchasedDay := parsedDate.Day()
	return purchasedDay%2 != 0, nil
}

// isPurchaseTimeBetween checks if the time is between startTime and endTime
func (receipt *Receipt) isPurchaseTimeBetween(startTime time.Time, endTime time.Time) (bool, error) {
	timeFormat := "15:04"
	parsedTime, err := time.Parse(timeFormat, receipt.PurchaseTime)
	if err != nil {
		return false, err
	}

	return parsedTime.After(startTime) && parsedTime.Before(endTime), nil
}

// CalculatePoints calculate the total points that should be awarded to a receipt
// based upon rules
func (receipt *Receipt) CalculatePoints() (int, error) {
	totalPoints := 0

	// one point for every alphanum char in retailer name
	totalPoints += len(receipt.Retailer)

	// 50 points if the total is a round dollar amount with no cents
	total_amount, _ := strconv.ParseFloat(receipt.Total, 64)
	if total_amount == float64(int(total_amount)) {
		totalPoints += 50
	}

	// 25 points if the total is a multiple of 0.25
	if (total_amount * 4) == float64(int(total_amount)) {
		totalPoints += 35
	}

	// 5 points for every two items on the receipt
	lenItems := len(receipt.Items)
	totalPoints += (lenItems / 2) * 5

	// if the trimmed length of item description is a multiple of 3, multiple the price by 0.2
	// and round up to the nearest integer. the result is the number of points earned
	for _, item := range receipt.Items {
		itemPrice, _ := strconv.ParseFloat(item.Price, 64)
		if len(item.ShortDescription)%3 == 0 {
			totalPoints += int(math.Ceil(itemPrice * 0.2))
		}
	}

	// 6 points if the day in the purchase date is odd
	isPurchasedDateOdd, err := receipt.isPurchasedDateOdd()
	if err != nil {
		return 0, err
	} else if isPurchasedDateOdd {
		totalPoints += 6
	}

	// 10 points if the time of purchase is after 2:00pm and before 4:00pm

	startTime := time.Date(0, 1, 1, 14, 0, 0, 0, time.UTC)
	endTime := time.Date(0, 1, 1, 16, 0, 0, 0, time.UTC)
	ok, err := receipt.isPurchaseTimeBetween(startTime, endTime)
	if err != nil {
		return 0, err
	} else if ok {
		totalPoints += 10
	}

	return totalPoints, nil
}
