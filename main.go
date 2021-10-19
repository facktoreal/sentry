package sentry

import (
	"errors"
	"log"
	"os"

	sentryGo "github.com/getsentry/sentry-go"
)

var(
	env = "production"
	release = ""
	dsn = ""
)

// Init sentry configuration
func Init(dsnStr string) error {
	// check if sentry DSN is set
	if len(dsnStr) == 0 {
		log.Println("No Sentry DNS provided, sentry reporting disabled")

		return nil
	}

	dsn = dsnStr

	if len(MayGetString("ENV")) > 0 {
		env = MayGetString("ENV")
	}

	// check if we provide release manually
	if len(MayGetString("RELEASE")) > 0 {
		release = MayGetString("RELEASE")
	}

	// check if we have release from appengine
	if len(MayGetString("GAE_VERSION")) > 0 {
		release = MayGetString("GAE_VERSION")
	}

	// Set release version
	if !MustPresent("RELEASE") {
		return errors.New("'RELEASE' must be set")
	}

	err := sentryGo.Init(sentryGo.ClientOptions{
		Dsn: dsn,
		Environment: env,
		Release: release,
	})

	if err != nil {
		return err
	}

	return nil
}

// CaptureError ...
func CaptureError(err error, tags map[string]string) {
	if len(dsn) == 0 {
		return
	}

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

// MayGetString ...
func MayGetString(key string) string {
	v := os.Getenv(key)
	if v == "" {
		return ""
	}

	return v
}