package fiberprom

import (
	"strconv"
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
			since := time.Since(start).Milliseconds()
			status := c.Response().StatusCode()
			httpRequestHistogramProm.WithLabelValues(
				c.Route().Path,
				utils.CopyString(c.Method()),
				strconv.Itoa(status)).
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
