package middleware

import (
	"context"
	"errors"
	"go-web-cli/internal/pkg/metrics"
	"go-web-cli/pkg/biz_name/domain"
	"net/http"
	"strconv"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	herrors "github.com/cloudwego/hertz/pkg/common/errors"
)

// Metrics record all endpoint metrics
func Metrics(ctx context.Context, c *app.RequestContext) {
	start := time.Now()
	// exec
	c.Next(ctx)

	errs := c.Errors
	var status int
	if len(errs) > 0 {
		err := errs[0]
		var e *domain.AppError
		if errors.As(err, &e) {
			// Get correct error code from domain.AppError type
			status = e.HTTPStatus
		}
		var fe *herrors.Error
		if errors.As(err, &fe) {
			// Get correct error code from herrors.Error type
			status = http.StatusInternalServerError
		}
	} else {
		status = c.Response.StatusCode()
	}

	statusCode := strconv.Itoa(status)

	mill := time.Since(start).Milliseconds()

	host := c.Request.Host()
	path := c.FullPath()
	method := c.Method()

	// set latency
	metrics.RequestLatency.WithLabelValues(string(host), string(method), path, statusCode).Observe(float64(mill))
	return
}
