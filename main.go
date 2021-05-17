//Package main here we create a server, that prints the current time in the browser.
//I am also instrumenting  my go server with prometheus
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

//function that keeps running the metrics increment, sleeps every 2 seconds
func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

var (
	//Setting up prometheus
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "CTM_APP_requests_served",
		Help: "The total number of processed events",
	})
)

var (
	httpPort string
)

func init() {
	//Setting up the server port
	httpPort = os.Getenv("SERVER_PORT")
	if len(httpPort) < 1 {
		httpPort = "8080"
	}
}

func main() {
	l := log.New(os.Stdout, "CTM-APPLICATION", log.LstdFlags|log.Lmicroseconds)
	if err := run(l); err != nil {
		l.Println(err)
		os.Exit(1)
	}

}

func run(l *log.Logger) error {

	//function that handles the root on http endpoint requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		//Reseting the http headers to prevent our app from being cached
		w.Header().Set("X-Accel-Expires", "0")
		w.Header().Set("Cache-Control", "no-cache, private, max-age=0")
		w.Header().Set("Expires", time.Unix(0, 0).Format(http.TimeFormat))
		w.Header().Set("Pragma", "no-cache")

		//Getting the current salutation/time and setting my http endpoint to show it on the page
		fmt.Fprint(w, saluteAndTime())
	})

	recordMetrics()
	//setting the metrics endpoint to show instrumentation data
	http.Handle("/metrics", promhttp.Handler())

	//starting the server, returns error or nil
	l.Println("starting the server on port:", httpPort)
	return http.ListenAndServe(":"+httpPort, nil)
}
