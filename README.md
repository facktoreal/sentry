# Environment Variables

## Sentry config

### `SENTRY_DSN`

Sentry DNS

### `ENV`

Current system environment. Supported options are:

* `DEV`
* `STG`
* `PROD`

### `RELEASE`

Current release version of backend. Can be set in CI/CD using:

```bash
sed -i app.yaml -e "s/__RELEASE__/$CI_COMMIT_TAG/"
```

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

    if err := sentry.Init(); err != nil {
        // handle error
    }

    // Activate echo middleware
    e.Use(sentry.Middleware())
}

func init() {
    if err = srv.Update(data); err != nil {
    	sentry.CaptureErrorAndWait(err, map[string]string{"category": "db"})

        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }
}
```
