package helpers

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func ConvertTimeToDuration(duration_string string) time.Duration {
	hours_replaced := strings.Replace(duration_string, ":", "h", 1)
	minutes_replaced := strings.Replace(hours_replaced, ":", "m", 1)
	seconds_replaced := fmt.Sprintf("%ss", minutes_replaced)
	duration, err := time.ParseDuration(seconds_replaced)
	if err != nil {
		log.Fatal(err)
	}
	return duration
}
