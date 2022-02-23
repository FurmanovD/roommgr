package healthcheck

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/FurmanovD/roommgr/internal/app/service"
	"github.com/FurmanovD/roommgr/pkg/api/v1"
)

// HealthChecker an EndpointHandler interface holder
type HealthChecker struct {
	service service.RoomManagerService
}

// NewHealthChecker instantiates a HealthChecker
func NewHealthChecker(svc service.RoomManagerService) *HealthChecker {
	return &HealthChecker{
		service: svc,
	}
}

// GetHealth is a GET /healthcheck handler
func (h *HealthChecker) GetHealth(c *gin.Context) {
	// TODO perform a service-specific healthcheck if required
	c.JSON(
		http.StatusOK,
		gin.H{
			"healthcheck": "Perfect",
			"time":        time.Now().UTC().Format(api.TimeFormat),
		},
	)
}
