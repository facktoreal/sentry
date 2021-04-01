package sentry

import (
	"errors"
	"log"
	"os"

	sentryGo "github.com/getsentry/sentry-go"
)

// Init sentry configuration
func Init() error {
	if !MustPresent("SENTRY_DSN") {
		return errors.New("'SENTRY_DSN' must be set")
	}

	if !MustPresent("ENV") {
		return errors.New("'ENV' must be set")
	}

	// Set release version
	if !MustPresent("RELEASE") {
		return errors.New("'RELEASE' must be set")
	}

	err := sentryGo.Init(sentryGo.ClientOptions{
		Dsn: MustGetString("SENTRY_DSN"),
		Environment: MustGetString("ENV"),
		Release: MustGetString("RELEASE"),
	})

	if err != nil {
		return err
	}

	return nil
}

// CaptureError ...
func CaptureError(err error, tags map[string]string) {
	sentryGo.WithScope(func(scope *sentryGo.Scope) {
		scope.SetContext("Request", tags)

		sentryGo.CaptureException(err)
	})
}

// MustPresent ...
func MustPresent(key string) bool {
	v := os.Getenv(key)
	if v == "" {
		return false
	}

	return true
}

// MustGetString ...
func MustGetString(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Panicf("%s environment variable not set.", key)
	}

	return v
}
