package https

import (
	"api-gateway/internal/https/handlers"
	authmiddleware "api-gateway/internal/https/middleware/AuthMiddleware"
	corsmiddleware "api-gateway/internal/https/middleware/corsMiddleware"
	"api-gateway/internal/pkg/service"
	"crypto/tls"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewGin(service *service.ServiceRepositoryClient) *http.Server {
	r := gin.Default()

	newHandlers := handlers.NewHandlers(service)
	r.Use(corsmiddleware.CorsMiddleware())
	r.Use(authmiddleware.Middleware())

	r.POST("/v1/workers/add-worker", newHandlers.AddWorkers)
	r.POST("/v1/workers/end-day", newHandlers.EndDay)
	r.POST("/v1/workers/load-blocks", newHandlers.LoadBlocks)
	r.POST("/v1/workers/paid_monthly/:worker_id", newHandlers.AddPaidMonthly)
	r.PUT("/v1/workers/update-worker/:worker_id", newHandlers.UpdateWorker)
	r.DELETE("/v1/workers/delete-worker/:worker_id", newHandlers.DeleteWorker)
	r.GET("/v1/workers/all-workers", newHandlers.AllWorkers)
	r.GET("/v1/workers/monthly-report", newHandlers.MonthlyReport)
	r.GET("/v1/workers/get-daily-production-id/:dailyproductionid", newHandlers.GetDailyProductionWorkersById)
	r.GET("/v1/workers/get-load-production-id/:loadproductionid", newHandlers.GetLoadProductionWorkersById)

	r.POST("/v1/users/login", newHandlers.Login)
	r.GET("/v1/check_me", newHandlers.CheckToken)

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		CurvePreferences:   []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	srv := &http.Server{
		Addr:      ":9000",
		Handler:   r,
		TLSConfig: tlsConfig,
	}

	return srv
}
