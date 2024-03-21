package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	timerForMalloc := time.NewTicker(time.Second * 5)
	memory := make(map[int][]byte)

	i := 0
	go func() {
		for range timerForMalloc.C {
			memory[i] = make([]byte, 1024)
		}
	}()

	http.Handle("/metric", promhttp.Handler())
	http.ListenAndServe(":8081", nil)
}
