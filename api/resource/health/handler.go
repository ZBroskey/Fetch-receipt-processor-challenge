package health

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Handler struct {
	logger zerolog.Logger
}

func NewHandler() *Handler {
	return &Handler{
		logger: log.With().Str("component", "health").Logger(),
	}
}

// @Summary Health Check
// @Description Check the health of the server
// @Produce json
// @Success 200
// @Router /health [get]

func (h *Handler) HealthCheck(c echo.Context) error {
	h.logger.Info().Msg("health check")

	return c.NoContent(http.StatusOK)
}