package services

import (
	"fmt"

	"github.com/newrelic/go-agent/v3/newrelic"
)

func StartNewrelic(licenseKey string) (*newrelic.Application, error) {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("Moonlogs"),
		newrelic.ConfigLicense(licenseKey),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)

	if err != nil {
		return nil, fmt.Errorf("failed starting newrelic: %w", err)
	}

	return app, nil
}
