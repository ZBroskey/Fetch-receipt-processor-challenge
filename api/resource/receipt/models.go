package receipt

type Receipt struct {
	retailer string	`json:"retailer"`
	purchasedDate string	`json:"purchasedDate"`
	purchasedTime string	`json:"purchasedTime"`
	items *Item	`json:"items"`
	total string	`json:"total"`
}

type Item []struct {
	shortDescription string	`json:"shortDescription"`
	price string	`json:"price"`
}