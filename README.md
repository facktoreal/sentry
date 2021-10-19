# Environment Variables

## Sentry config

### `SENTRY_DSN`

Sentry DNS

### `RELEASE`

Current release version of backend. Can be set as ENV `RELEASE`

### Usage

Current release version of backend. Can be set in CI/CD using:

```go
package main

import (
    "github.com/labstack/echo/v4"
    "yn.ee/facktoreal/sentry"
)

func main()  {
    e := echo.New()

    if err := sentry.Init("%SENTRY_DSN%"); err != nil {
        // handle error
    }

    // Activate echo middleware
    e.Use(sentry.Middleware())
}

func init() {
    if err = srv.Update(data); err != nil {
    	sentry.CaptureErrorw(err, map[string]string{"category": "db"})

        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }
}
```
