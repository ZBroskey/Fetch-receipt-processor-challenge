package router

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	// "/fetch-receipt-processor-challenge/api/resource/health"
)

func New(l *zerolog.Logger) {
	e := echo.New()

	// e.GET("/health", health.HealthCheck)

	// rpApi := e.Group("/api/v1/receipts")
	// rpApi.GET("/:id/points", receipt.GetPoints)
	// rpApi.POST("/process", receipt.ProcessReceipt)

	e.Logger.Fatal(e.Start(":8080"))
}