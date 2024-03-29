package fiberprom

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	httpRequestHistogramProm = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_request_histogram",
		Help:    "Histogram of the http request duration.",
		Buckets: prometheus.LinearBuckets(1, 1, 10),
		// Buckets: []float64{
		// 	0.000000001, // 1ns
		// 	0.000000002,
		// 	0.000000005,
		// 	0.00000001, // 10ns
		// 	0.00000002,
		// 	0.00000005,
		// 	0.0000001, // 100ns
		// 	0.0000002,
		// 	0.0000005,
		// 	0.000001, // 1µs
		// 	0.000002,
		// 	0.000005,
		// 	0.00001, // 10µs
		// 	0.00002,
		// 	0.00005,
		// 	0.0001, // 100µs
		// 	0.0002,
		// 	0.0005,
		// 	0.001, // 1ms
		// 	0.002,
		// 	0.005,
		// 	0.01, // 10ms
		// 	0.02,
		// 	0.05,
		// 	0.1, // 100 ms
		// 	0.2,
		// 	0.5,
		// 	1.0, // 1s
		// 	2.0,
		// 	5.0,
		// 	10.0, // 10s
		// 	15.0,
		// 	20.0,
		// 	30.0,
		// },
	}, []string{"path", "method", "status"})
)

// CopyString copies a string to make it immutable

func New() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Path() == "/metrics" {
			return c.Next()
		}

		start := time.Now()
		defer func() {
			since := time.Since(start).Microseconds()
			status := c.Response().StatusCode()
			httpRequestHistogramProm.WithLabelValues(
				c.Route().Path,
				utils.CopyString(c.Method()),
				http.StatusText(status)).
				Observe(float64(since))
		}()

		// serve the request to the next middleware
		if err := c.Next(); err != nil {
			// invokes the registered HTTP error handler
			// to get the correct response status code
			_ = c.App().Config().ErrorHandler(c, err)
		}

		return nil
	}
}
