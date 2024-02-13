package index

import (
	"fmt"
	"time"
)

func Current() string {
	now := time.Now()
	month := now.Month()
	year := now.Year()

	return fmt.Sprintf("%d-%d", month, year)
}
