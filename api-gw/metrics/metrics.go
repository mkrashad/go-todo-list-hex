package metrics

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)


var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number click events on users enpoint",
	})
)

func PrometheusMetrics(r *gin.Engine) {
	// Setup Prometheus
	p := ginprometheus.NewPrometheus("gin")

	r.Use(func(c *gin.Context) {
		path := strings.Split(c.Request.URL.String(), "/")
		if path[len(path)-1] == "tasks" {
			opsProcessed.Inc()
		}
	})

	p.Use(r)
}
