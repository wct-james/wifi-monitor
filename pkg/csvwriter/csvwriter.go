package csvwriter

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
	"wifi_monitor/pkg/speedtest"
)

const (
	FILENAME string = "/home/wjames/Documents/Projects/wifi-monitor/data/wifi-monitor.csv"
)

var CSVHeaders []string = []string{
	"time",
	"server",
	"latency_ms",
	"download_mb",
	"upload_mb",
	"error",
}

func CSVAppend(results []speedtest.SpeedTestResult) error {
	exists := true
	if _, err := os.Stat(FILENAME); os.IsNotExist(err) {
		exists = false
	}

	// Open file in appropriate mode
	file, err := os.OpenFile(FILENAME, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// For new files, write headers first
	if !exists {
		if err := writer.Write(CSVHeaders); err != nil {
			return fmt.Errorf("failed to write headers: %v", err)
		}
	}

	records := make([][]string, len(results))

	for i, res := range results {
		records[i] = []string{
			res.ResultTime.Format(time.RFC3339),
			res.Server,
			fmt.Sprintf("%d", res.Latency),
			fmt.Sprintf("%.2f", res.Download),
			fmt.Sprintf("%.2f", res.Upload),
			res.Error,
		}
	}

	// Write records
	if err := writer.WriteAll(records); err != nil {
		return fmt.Errorf("failed to write records: %v", err)
	}

	return nil
}
