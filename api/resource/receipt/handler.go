package receipt

import (
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/ZBroskey/Fetch-receipt-processor-challenge/api/resource/common/tools"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)
type Handler struct {
	logger zerolog.Logger
}

func NewHandler() *Handler {
	return &Handler{
		logger: log.With().Str("component", "receipt").Logger(),
	}
}

func getReceiptPoints(receipt *Receipt) int {
	points := 0

	// receipt.Retailer: Points for Alphanumeric Characters
	name := strings.ReplaceAll(receipt.Retailer, " ", "")
	for range name {
		points += 1
	}

	// receipt.Total: Points for Round Number
	total, err := strconv.ParseFloat(receipt.Total, 64)
	
	if err != nil {
		return 0
	}

	if total == math.Trunc(total) {
		points += 50
	}

	// receipt.Total: for Multiple of 0.25
	if math.Mod(total, 0.25) == 0 {
		points += 25
	}
	
	// receipt.Items: Points for Even Number of Items
	points += len(receipt.Items) / 2

	// receipt.Items: Points for Trimmed Description
	for _, item := range receipt.Items {
		if len(strings.TrimSpace(item.ShortDescription)) / 3 > 0 {
			price, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				continue
			}
			points += int(math.Ceil(price * 0.2))
		}
	}

	// receipt.PurchasedDate: Points for Odd Date
	date, err := strconv.Atoi(receipt.PurchasedDate[len(receipt.PurchasedDate) - 2:])
	if err != nil {
		return 0
	}
	if date % 2 != 0 {
		points += 6
	}

	// receipt.PurchasedTime: Points for Purchase Between 2pm and 4pm
	time, err := strconv.ParseFloat(strings.Replace(receipt.PurchasedTime, ":", ".", 1), 64)
	if err != nil {
		return 0
	}
	if time >= 14 && time <= 16 {
		points += 10
	}

	return points
}

// @Summary					Get points
// @Description			Get the points for an existing receipt
// @Produce					json
// @Param						:id path string
// @Success 				200
// @Failure 				400
// @Failure 				404
// @Router 					/api/v1/receipts/{id}/points [get]
func (h *Handler) GetPoints(c echo.Context) error {
	h.logger.Info().Msg("Get points")

	receiptId := c.Param("id")

	if !tools.IsValidId(receiptId) {
		return echo.NewHTTPError(http.StatusBadRequest, ErrInvalidReceipt)
	}

	receipt, err := Find(c.Request().Context(), receiptId)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	// Point Calculation
	points := getReceiptPoints(receipt)

	return c.JSON(http.StatusOK, points)
}

// @Summary					Process receipt
// @Description			Process a receipt and saves in temporary storage
// @Accept					json
// @Produce					json
// @Param						receipt body Receipt
// @Success 				201
// @Failure 				400
// @Router 					/api/v1/receipts/process [post]
func (h *Handler) ProcessReceipt(c echo.Context) error {
	h.logger.Info().Msg("Process receipt")

	var receipt Receipt
	if err := c.Bind(&receipt); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}

	if valid, err := ValidateReceipt(receipt); !valid {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	receiptId := tools.RandomId()

	if err := Save(c.Request().Context(), receiptId, receipt); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to save receipt")
	}

	return c.JSON(http.StatusCreated, receiptId)
}