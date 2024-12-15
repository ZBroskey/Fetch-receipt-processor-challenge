package receipt

import (
	"regexp"
	"strconv"
)

type Receipt struct {
	Retailer string	`json:"retailer"`
	PurchasedDate string	`json:"purchasedDate"`
	PurchasedTime string	`json:"purchasedTime"`
	Items []Item	`json:"items"`
	Total string	`json:"total"`
}

type Item struct {
	ShortDescription string	`json:"shortDescription"`
	Price string	`json:"price"`
}

// validateRequired validates the required fields of a receipt
// - Retailer is required
// - PurchasedDate is required
// - PurchasedTime is required
// - Items are required
// - Total is required
func validateRequired(receipt Receipt) (bool, error) {
	var err error = nil
	switch {
	case receipt.Retailer == "":
		err = ErrRetailerRequired
	case receipt.PurchasedDate == "":
		err = ErrPurchasedDateRequired
	case receipt.PurchasedTime == "":
		err = ErrPurchasedTimeRequired
	case len(receipt.Items) == 0:
		err = ErrItemsRequired
	case receipt.Total == "":
		err = ErrTotalRequired
	default:
		return true, nil
	}

	return false, err
}

// ValidateReceipt validates a receipt
// - PurchasedDate must be in the format YYYY-MM-DD
// - PurchasedTime must be in the format HH:MM
// - All Prices must be in the format DDD.DD
// - Total must be the sum of all item prices
func ValidateReceipt(receipt Receipt) (bool, error) {
	if valid, err := validateRequired(receipt); !valid {
		return false, err
	}
	
	datePattern := `^\d{4}-\d{2}-\d{2}$`
	timePattern := `^\d{2}:\d{2}$`
	pricePattern := `^\d+\.\d{2}$`

	matched, _ := regexp.MatchString(datePattern, receipt.PurchasedDate)
	if !matched {
		return false, ErrInvalidPurchasedDate
	}

	matched, _ = regexp.MatchString(timePattern, receipt.PurchasedTime)
	if !matched {
		return false, ErrInvalidPurchasedTime
	}

	matched, _ = regexp.MatchString(pricePattern, receipt.Total)
	if !matched {
		return false, ErrInvalidTotal
	}

	total := 0.0

	for _, item := range receipt.Items {
		matched, _ = regexp.MatchString(pricePattern, item.Price)
		if !matched {
			return false, ErrInvalidPrice
		}

		price, _ := strconv.ParseFloat(item.Price, 64)
		total += price
	}

	receiptTotal, _ := strconv.ParseFloat(receipt.Total, 64)

	if total != receiptTotal {
		return false, ErrTotalDoesNotCompute
	}

	return true, nil
}