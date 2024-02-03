package timerange

import (
	"fmt"
	"net/http"
	"time"
)

const (
	layout = "2006-01-02T15:04"
)

func parseWithLocation(timeStr string, tzName string) (*time.Time, error) {
	if timeStr == "" {
		return nil, nil
	}

	if tzName == "" {
		tzName = "UTC"
	}

	location, err := time.LoadLocation(tzName)
	if err != nil {
		return nil, fmt.Errorf("loading location: %w", err)
	}

	parsedTime, err := time.ParseInLocation(layout, timeStr, location)
	if err != nil {
		return nil, fmt.Errorf("parsing time: %w", err)
	}
	return &parsedTime, nil
}

func Parse(r *http.Request) (*time.Time, *time.Time, error) {
	fromStr := r.URL.Query().Get("from")
	toStr := r.URL.Query().Get("to")
	tzStr := r.URL.Query().Get("tz")

	from, err := parseWithLocation(fromStr, tzStr)
	if err != nil {
		return nil, nil, err
	}

	to, err := parseWithLocation(toStr, tzStr)
	if err != nil {
		return nil, nil, err
	}

	return from, to, nil
}
