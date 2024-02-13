package commands

import "time"

func nowWithPtr() *time.Time {
	t := time.Now()
	return &t
}
