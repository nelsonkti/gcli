/**
** @创建时间 : 2022/3/26 11:36
** @作者 : fzy
 */
package middleware

import (
	"demod/lib/logger"
	"github.com/go-kratos/aegis/ratelimit"
	"github.com/go-kratos/aegis/ratelimit/bbr"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func RateLimiterMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		limiter := bbr.NewLimiter(
			bbr.WithWindow(5*time.Second),
			bbr.WithBucket(1),
			bbr.WithCPUThreshold(1))

		allow, err := limiter.Allow()

		if err != nil {
			logger.Sugar.Error(err)
			return c.JSON(http.StatusUnauthorized, err)
		}

		defer func() {
			if err == nil {
				allow(ratelimit.DoneInfo{})
			}
		}()

		return next(c)
	}
}