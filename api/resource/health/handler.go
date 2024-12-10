package health

import (
	"net/http"
)

// @Summary Health Check
// @Description Check the health of the service
// @ID health-check
// @Produce json
// @Success 200
// @Router /health [get]

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}