package utils

import (
	"fmt"
	"log"
	"time"
)

func formatDuration(d time.Duration) string {
	if d < time.Microsecond {
		return fmt.Sprintf("%d ns", d.Nanoseconds())
	} else if d < time.Millisecond {
		return fmt.Sprintf("%.2f µs", float64(d.Nanoseconds())/1000)
	} else if d < time.Second {
		return fmt.Sprintf("%.2f ms", float64(d.Nanoseconds())/1000000)
	}
	return fmt.Sprintf("%.2f s", d.Seconds())
}

func LogWithEmoji(emoji, msg string, start time.Time) {
	log.Printf("%s %s in %s", emoji, msg, formatDuration(time.Since(start)))
}
