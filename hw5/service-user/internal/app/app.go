package app

import (
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

const defaultAddr = ":8080"
const defaultMetricsPath = "/metrics"

func getEnv(key, defaultValue string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}

	return val
}

type App struct {
	srv     *gin.Engine
	addr    string
	metrics string
}

func NewFromEnv() *App {
	srv := gin.Default()
	addr := getEnv("SERVICE_URL", defaultAddr)
	metricsPath := getEnv("SERVICE_METRICS_PATH", defaultMetricsPath)

	return &App{srv: srv, addr: addr, metrics: metricsPath}
}

func (a *App) Run() error {
	prometheus.MustRegister(requestDuration)
	a.srv.GET(a.metrics, metricHandler())

	return a.srv.Run(a.addr)
}

func (a *App) GET(relativePath string, handler Handler) {
	a.srv.GET(relativePath, a.handlerWrapper(relativePath, "GET", handler))
}

func (a *App) POST(relativePath string, handler Handler) {
	a.srv.POST(relativePath, a.handlerWrapper(relativePath, "POST", handler))
}

func (a *App) DELETE(relativePath string, handler Handler) {
	a.srv.DELETE(relativePath, a.handlerWrapper(relativePath, "DELETE", handler))
}

func (a *App) PATCH(relativePath string, handler Handler) {
	a.srv.PATCH(relativePath, a.handlerWrapper(relativePath, "PATCH", handler))
}

func (a *App) handlerWrapper(endpoint string, method string, handler Handler) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		start := time.Now()
		defer func() {
			labeled := requestDuration.WithLabelValues(endpoint, method, strconv.Itoa(ctx.Writer.Status()))
			labeled.Observe(time.Since(start).Seconds())
		}()

		request := &Request{ctx: ctx}
		response, err := handler(ctx, request)

		switch err {
		case nil:
			ctx.JSON(200, gin.H{
				"result": gin.H(response),
			})
		case ValidationError:
			ctx.JSON(400, gin.H{
				"error": err.Error(),
			})
		case NotFound:
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
		case PermissionDenied:
			ctx.JSON(403, gin.H{
				"error": err.Error(),
			})
		default:
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
		}
	}
}
