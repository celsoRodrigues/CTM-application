package main

import "fmt"

func saluteAndTime() string {
	ts := fmt.Sprintf(Salute() + "\n" + myTime())
	return ts
}
