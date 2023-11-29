package sentry

import (
	sentryEcho "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo/v4"
)

func Middleware(options sentryEcho.Options) echo.MiddlewareFunc {
	return sentryEcho.New(options)
}
