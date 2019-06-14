package sentry

import (
	"errors"
	"fmt"
	"runtime/debug"

	"github.com/getsentry/raven-go"
	"github.com/labstack/echo/v4"
)

// Sentry struct holding the raven client and some of its configs
type Sentry struct {
	RavenClient *raven.Client
	Tags        map[string]string
}

// TagsFunc given a request context, extract some additional tags and return
// them as map[string]string as required by the raven client.
type TagsFunc func(c echo.Context) map[string]string

var (
	sentry   = &Sentry{}
	tagsFunc TagsFunc
)

// SetTags sets any other additional tags to be captured by Sentry.
// Tags can be extracted from the current request context
// or just static tags, e.g. tags["app_version"] = appVersion.
func SetTags(fn TagsFunc) {
	tagsFunc = fn
}

// Middleware returns an echo middleware which recovers from panics anywhere in the chain
// and logs to the sentry server specified in DSN.
func Middleware() echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				sentry.Tags = map[string]string{
					"endpoint": c.Request().URL.String(),
				}

				if rval := recover(); rval != nil {
					debug.PrintStack()

					errorMsg := fmt.Sprint(rval)
					err := errors.New(errorMsg)

					stacktrace := raven.NewException(err, raven.NewStacktrace(2, 3, nil))

					httpContext := raven.NewHttp(c.Request())

					// extract tags
					if tagsFunc != nil {
						sentry.Tags = tagsFunc(c)
					}

					// contruct the raven packet to be sent
					packet := raven.NewPacket(errorMsg, stacktrace, httpContext)

					// capture it and send.
					raven.Capture(packet, sentry.Tags)

					// register the error back to echo.Context
					c.Error(err)
				}
			}()

			raven.SetHttpContext(raven.NewHttp(c.Request()))

			return h(c)
		}
	}
}
