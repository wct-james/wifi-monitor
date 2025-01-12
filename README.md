## WiFi Monitor

Run this in a cron job to test your WiFi speed regularly 

## Usage

1. Build the binary:
`go build -o ./build/wifi-monitor cmd/main.go`

2. Add the executable to a cron job:

e.g. Every 5 minutes

Run: `crontab -e`
Then insert this line to the file
```
*/5 * * * * /path/to/bin/wifi-monitor >> /home/user/logs/wifi-monitor.log 2>&1
```
