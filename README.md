# Fetch-receipt-processor-challenge

## Summary

This is my response to the [Fetch Receipt Processor Challenge](https://github.com/fetch-rewards/receipt-processor-challenge)

## Getting Started

- Clone the repository

  - `git clone https://github.com/ZBroskey/Fetch-receipt-processor-challenge.git`

- Change to Fetch-receipt-processor-challenge/cmd/api directory

  - `cd Fetch-receipt-processor-challenge/cmd/api`

- Run main.go
  - `go run .`
  - Default port is 8081, however, a specific port can be chosen with the `--port ####` flag

## Additional

An example receipt is provided with the following structure:

    `"7fb1377b-b223-49d9-a31a-5a02701dd310": {
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
    }`
