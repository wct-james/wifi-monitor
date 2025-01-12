package main

import (
	"wifi_monitor/pkg/csvwriter"
	"wifi_monitor/pkg/speedtest"
)

func main() {
	results := speedtest.SpeedTest()
	if err := csvwriter.CSVAppend(results); err != nil {
		panic("could not write file")
	}
}
