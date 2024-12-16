package receipt

import "context"

type Repository interface {
	Find(ctx context.Context, id string) (*Receipt, error)
	Save(ctx context.Context, id string, receipt Receipt) error
}

var receipts = map[string]Receipt{
	"7fb1377b-b223-49d9-a31a-5a02701dd310": {
		Retailer: "M&M Corner Market",
		PurchaseDate: "2022-03-20",
		PurchaseTime: "14:33",
		Items: []Item{
			{
				ShortDescription: "Gatorade",
				Price: "2.25",
			},{
				ShortDescription: "Gatorade",
				Price: "2.25",
			},{
				ShortDescription: "Gatorade",
				Price: "2.25",
			},{
				ShortDescription: "Gatorade",
				Price: "2.25",
			},
		},
		Total: "9.00",
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