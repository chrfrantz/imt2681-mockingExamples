package main

import (
	"time"
	"fmt"
	"strconv"
)

func Benchmark(start time.Time) {
	duration := time.Now().Nanosecond()/1e6 - start.Nanosecond()/1e6
	fmt.Println("Duration: " + strconv.Itoa(duration) + "ms")
}
