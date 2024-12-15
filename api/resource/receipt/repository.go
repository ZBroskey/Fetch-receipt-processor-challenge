package receipt

import "context"

type Repository interface {
	Find(ctx context.Context, id string) (*Receipt, error)
	Save(ctx context.Context, id string, receipt Receipt) error
}

var receipts = map[string]Receipt{
	"7fb1377b-b223-49d9-a31a-5a02701dd310": {
			Retailer:     "Target",
			PurchasedDate: "2022-01-01",
			PurchasedTime: "13:01",
			Items: []Item{
					{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
					{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
					{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
					{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
					{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
			},
			Total: "35.35",
	},
}

func Find(ctx context.Context, id string) (*Receipt, error) {
	receipt, exists := receipts[id]
	if !exists {
		return nil, ErrReceiptNotFound
	}
	return &receipt, nil
}

func Save(ctx context.Context, id string, receipt Receipt) error {
	receipts[id] = receipt
	return nil
}