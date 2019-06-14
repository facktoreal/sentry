package sentry

import (
	"errors"

	"github.com/getsentry/raven-go"
	"yn.ee/facktoreal/env"
)

// Init sentry configuration
func Init() error {
	if !env.MustPresent("SENTRY_DSN") {
		return errors.New("'SENTRY_DSN' must be set")
	}

	if !env.MustPresent("ENV") {
		return errors.New("'ENV' must be set")
	}

	// Set release version
	if !env.MustPresent("RELEASE") {
		return errors.New("'RELEASE' must be set")
	}

	// Enable sentry integration
	if err := raven.SetDSN(env.MustGetString("SENTRY_DSN")); err != nil {
		return err
	}

	// Configure sentry
	raven.SetRelease(env.MustGetString("RELEASE"))
	raven.SetEnvironment(env.MustGetString("ENV"))

	return nil
}

// CaptureErrorAndWait ...
func CaptureErrorAndWait(err error, tags map[string]string, interfaces ...raven.Interface) string {
	return raven.CaptureErrorAndWait(err, tags, interfaces...)
}
