package main

import (
	"github.com/ZBroskey/Fetch-receipt-processor-challenge/api/resource/health"
	"github.com/ZBroskey/Fetch-receipt-processor-challenge/api/resource/receipt"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func init() {
	log.Info().Msg("init started")
}

// @Title						Receipt Processor API
// @Version					1.0
// @Description			The Receipt Processor API provides endpoints to process receipts and get points from receipts.
//
// @Host						localhost:8081
//
// @Consumes				application/json
// @Produces				application/json
func main() {
	log.Info().Msg("setup started")

	e := echo.New()

	healthHandler := health.NewHandler()
	receiptHandler := receipt.NewHandler()

	// /health
	e.GET("/health", healthHandler.HealthCheck)

	rpApi := e.Group("/api/v1/receipts")

	// /api/v1/receipts
	rpApi.GET("/:id/points", receiptHandler.GetPoints)
	rpApi.POST("/process", receiptHandler.ProcessReceipt)

	e.Logger.Fatal(e.Start(":8081"))
}