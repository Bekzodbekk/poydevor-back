package https

import (
	"api-gateway/internal/https/handlers"
	corsmiddleware "api-gateway/internal/https/middleware/corsMiddleware"
	"api-gateway/internal/pkg/service"

	"github.com/gin-gonic/gin"
)

func NewGin(service *service.ServiceRepositoryClient) *gin.Engine {
	r := gin.Default()

	newHandlers := handlers.NewHandlers(service)
	r.Use(corsmiddleware.CorsMiddleware())

	r.POST("/v1/workers/add-worker", newHandlers.AddWorkers)
	r.POST("/v1/workers/end-day", newHandlers.EndDay)
	r.POST("/v1/workers/load-blocks", newHandlers.LoadBlocks)
	r.GET("/v1/workers/all-workers", newHandlers.AllWorkers)

	return r
}
