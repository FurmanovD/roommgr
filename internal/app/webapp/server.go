package webapp

import (
	"errors"

	"github.com/gin-gonic/gin"

	"github.com/FurmanovD/roommgr/internal/app/service"
	"github.com/FurmanovD/roommgr/pkg/ginserver"
)

// webServer implements the WebServer interface
type webServer struct {
	roommgrService service.RoomManagerService
	ginServer      ginserver.GinServer

	routerGroupRoot  *gin.RouterGroup
	routerGroupAPIV1 *gin.RouterGroup
}

// NewServer creates a WebServer interface instance
func NewServer(roommgrService service.RoomManagerService) WebServer {
	srv := ginserver.NewGinServer(true)
	return &webServer{
		roommgrService:   roommgrService,
		ginServer:        srv,
		routerGroupRoot:  srv.Engine().Group(""),
		routerGroupAPIV1: srv.Engine().Group("/api").Group("/v1"),
	}
}

// ListenAndServe starts routes serve on the given port
func (s *webServer) ListenAndServe(port int) error {
	if s.ginServer == nil {
		return errors.New("underlying gin server has not been initialized")
	}

	return s.ginServer.ListenAndServe(port)
}
