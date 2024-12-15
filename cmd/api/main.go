package main

import (
	"context"

	"github.com/ZBroskey/Fetch-receipt-processor-challenge/api/resource/health"
	"github.com/ZBroskey/Fetch-receipt-processor-challenge/api/resource/receipt"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

var ctx = context.Background()


func init() {
	log.Info().Msg("init started")
}

// @title							Receipt Processor
// @version						1.0
// @description				The Receipt Processor API provides endpoints to process receipts
//										and calculate points of existing receipts.
//
// @contact.name			Zachary Broskey
// @contact.email			zbroskey@me.com
//
// @host							localhost:8080
func main() {
	log.Info().Msg("setup started")

	e := echo.New()

	healthHandler := health.NewHandler()
	receiptHandler := receipt.NewHandler()

	e.GET("/health", healthHandler.HealthCheck)

	rpApi := e.Group("/api/v1/receipts")
	rpApi.GET("/:id/points", receiptHandler.GetPoints)
	rpApi.POST("/process", receiptHandler.ProcessReceipt)

	e.Logger.Fatal(e.Start(":8080"))
}