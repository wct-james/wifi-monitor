package speedtest

import (
	"log/slog"
	"time"

	"github.com/showwin/speedtest-go/speedtest"
)

type SpeedTestResult struct {
	ResultTime time.Time
	Server     string
	Latency    int64
	Download   float64
	Upload     float64
	Error      string
}

func SpeedTestError(err error) SpeedTestResult {
	return SpeedTestResult{
		time.Now(),
		"",
		0, 0.0, 0.0,
		err.Error(),
	}
}

func SpeedTest() []SpeedTestResult {
	result := make([]SpeedTestResult, 0)
	slog.Info("starting speed test")

	serverList, err := speedtest.FetchServers()
	if err != nil {
		slog.Error("error fetching server list", "error", err)
		return append(result, SpeedTestError(err))
	}

	// Find closest server
	targets, err := serverList.FindServer([]int{})
	if err != nil {
		slog.Error("error finding server", "error", err)
		return append(result, SpeedTestError(err))
	}

	for _, server := range targets {
		slog.Info("testing against", "name", server.Name, "country", server.Country)

		if server.Name != "London" {
			break
		}

		// Test download
		err = server.DownloadTest()
		if err != nil {
			slog.Error("error testing download", "error", err)
			result = append(result, SpeedTestError(err))
			continue
		}

		// Test upload
		err = server.UploadTest()
		if err != nil {
			slog.Error("error testing upload", "error", err)
			result = append(result, SpeedTestError(err))
			continue
		}

		result = append(result, SpeedTestResult{
			time.Now(),
			server.Name,
			server.Latency.Microseconds(),
			server.DLSpeed.Mbps(),
			server.ULSpeed.Mbps(),
			"",
		})
	}

	slog.Info("speed test completed")

	return result
}
