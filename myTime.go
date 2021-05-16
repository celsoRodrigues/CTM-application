package main

import "time"

//myTime function returns the current time
func myTime() string {
	t := time.Now()
	return t.String()
}
