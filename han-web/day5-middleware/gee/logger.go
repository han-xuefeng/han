package gee

import (
	"fmt"
	"time"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		// Start timer
		t := time.Now()
		// Process request
		c.Next()
		// Calculate resolution time
		fmt.Println("$$$$$$$$$$$", t)
	}
}