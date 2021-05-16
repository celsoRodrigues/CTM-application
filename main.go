//Package main here we create our simple server, that prints the current time in the browser.
//I am also instrumenting  my go server with prometheus
package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})
)

var httpPort string

func init() {
	httpPort = os.Getenv("HTTP_PORT")
	if len(httpPort) < 1 {
		httpPort = "80"
	}
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		//Reseting the http headers to prevent our app from being cached
		w.Header().Set("X-Accel-Expires", "0")
		w.Header().Set("Cache-Control", "no-cache, private, max-age=0")
		w.Header().Set("Expires", time.Unix(0, 0).Format(http.TimeFormat))
		w.Header().Set("Pragma", "no-cache")

		t := time.Now()
		fmt.Fprint(w, "Hello\n", t.String())
	})

	recordMetrics()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(httpPort, nil)
}
