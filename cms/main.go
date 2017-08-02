package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/geneseeq/authorize-system/cms/action"
	"github.com/geneseeq/authorize-system/cms/grouping"
	"github.com/geneseeq/authorize-system/cms/usering"

	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	defaultPort = "8080"
)

func main() {
	var (
		addr     = envString("PORT", defaultPort)
		httpAddr = flag.String("http.addr", ":"+addr, "HTTP listen address")
	)

	flag.Parse()

	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	var (
		// set db and collection
		groups = action.NewGroupDBRepository("test", "groups")
		users  = action.NewUserDBRepository("test", "user")
	)

	fieldKeys := []string{"method"}

	var gs grouping.Service
	gs = grouping.NewService(groups)
	gs = grouping.NewLoggingService(log.With(logger, "component", "grouping"), gs)
	gs = grouping.NewInstrumentingService(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "grouping_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "grouping_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		gs,
	)

	var us usering.Service
	us = usering.NewService(users)
	us = usering.NewLoggingService(log.With(logger, "component", "usering"), us)
	us = usering.NewInstrumentingService(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "usering_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "usering_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		us,
	)

	httpLogger := log.With(logger, "component", "http")

	mux := http.NewServeMux()

	mux.Handle("/grouping/v1/", grouping.MakeHandler(gs, httpLogger))
	mux.Handle("/usering/v1/", usering.MakeHandler(us, httpLogger))

	http.Handle("/", accessControl(mux))
	http.Handle("/metrics", promhttp.Handler())

	errs := make(chan error, 2)
	go func() {
		logger.Log("transport", "http", "address", *httpAddr, "msg", "listening")
		errs <- http.ListenAndServe(*httpAddr, nil)
	}()
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	logger.Log("terminated", <-errs)
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
