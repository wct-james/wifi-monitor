package speedtest

import (
	"fmt"
	"log"
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
	fmt.Println("Starting Speed Test...")

	serverList, err := speedtest.FetchServers()
	if err != nil {
		log.Printf("Error fetching server list: %v\n", err)
		return append(result, SpeedTestError(err))
	}

	// Find closest server
	targets, err := serverList.FindServer([]int{})
	if err != nil {
		log.Printf("Error finding server: %v\n", err)
		return append(result, SpeedTestError(err))
	}

	for _, server := range targets {
		fmt.Printf("Testing against %s (%s)...\n", server.Name, server.Country)

		if server.Name != "London" {
			break
		}

		// Test download
		err = server.DownloadTest()
		if err != nil {
			log.Printf("Error testing download: %v\n", err)
			result = append(result, SpeedTestError(err))
			continue
		}

		// Test upload
		err = server.UploadTest()
		if err != nil {
			log.Printf("Error testing upload: %v\n", err)
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

	return result
}
