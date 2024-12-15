package receipt

import (
	"errors"
)

var (
	ErrInvalidReceipt = errors.New("invalid receipt id")

	ErrReceiptNotFound = errors.New("receipt id not found")

	ErrRetailerRequired = errors.New("retailer is required")

	ErrPurchasedDateRequired = errors.New("purchased date is required")

	ErrPurchasedTimeRequired = errors.New("purchased time is required")

	ErrItemsRequired = errors.New("items are required")

	ErrTotalRequired = errors.New("total is required")

	ErrInvalidPurchasedDate = errors.New("purchased date is invalid")

	ErrInvalidPurchasedTime = errors.New("purchased time is invalid")

	ErrInvalidPrice = errors.New("price is invalid")

	ErrInvalidTotal = errors.New("total is invalid")

	ErrTotalDoesNotCompute = errors.New("item prices do not add up to total")
)