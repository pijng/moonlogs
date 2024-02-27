package services

import (
	"fmt"
	"os"
	"runtime"

	"github.com/grafana/pyroscope-go"
)

func StartPyroscope(address string) error {
	if address == "" {
		return fmt.Errorf("pyroscope address is empty, specify it in config via \"pyroscope_address\": <address>")
	}

	runtime.SetMutexProfileFraction(5)
	runtime.SetBlockProfileRate(5)

	_, err := pyroscope.Start(pyroscope.Config{
		ApplicationName: "moonlogs",
		ServerAddress:   address,
		// Consider adding feature to enable stdout logs
		Logger: nil,
		Tags:   map[string]string{"hostname": os.Getenv("HOSTNAME")},
		ProfileTypes: []pyroscope.ProfileType{
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,
			pyroscope.ProfileGoroutines,
			pyroscope.ProfileMutexCount,
			pyroscope.ProfileMutexDuration,
			pyroscope.ProfileBlockCount,
			pyroscope.ProfileBlockDuration,
		},
	})

	return err
}
