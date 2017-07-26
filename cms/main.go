package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
)

const (
	defaultPort              = "8080"
	defaultRoutingServiceURL = "http://localhost:7878"
)

func main() {
	var (
		addr  = envString("PORT", defaultPort)
		rsurl = envString("ROUTINGSERVICE_URL", defaultRoutingServiceURL)

		httpAddr          = flag.String("http.addr", ":"+addr, "HTTP listen address")
		routingServiceURL = flag.String("service.routing", rsurl, "routing service URL")
		//ctx为空
		ctx = context.Background()
	)

	flag.Parse()

	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	var (
		users = action.NewUserRepository()
		// locations      = inmem.NewLocationRepository()
		// voyages        = inmem.NewVoyageRepository()
		// handlingEvents = inmem.NewHandlingEventRepository()
	)

	// // Configure some questionable dependencies.
	// var (
	// 	handlingEventFactory = cargo.HandlingEventFactory{
	// 		CargoRepository:    cargos,
	// 		VoyageRepository:   voyages,
	// 		LocationRepository: locations,
	// 	}
	// 	handlingEventHandler = handling.NewEventHandler(
	// 		inspection.NewService(cargos, handlingEvents, nil),
	// 	)
	// )

	// // Facilitate testing by adding some cargos.
	// storeTestData(cargos)

	fieldKeys := []string{"method"}

	// var rs routing.Service
	// rs = routing.NewProxyingMiddleware(ctx, *routingServiceURL)(rs)

	var bs usering.Service
	bs = usering.NewService(users)
	bs = usering.NewLoggingService(log.With(logger, "component", "usering"), bs)
	bs = usering.NewInstrumentingService(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "booking_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "booking_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		bs,
	)

	httpLogger := log.With(logger, "component", "http")

	mux := http.NewServeMux()

	mux.Handle("/user/v1/", usering.MakeHandler(bs, httpLogger))

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

// func storeTestData(r cargo.Repository) {
// 	test1 := cargo.New("FTL456", cargo.RouteSpecification{
// 		Origin:          location.AUMEL,
// 		Destination:     location.SESTO,
// 		ArrivalDeadline: time.Now().AddDate(0, 0, 7),
// 	})
// 	if err := r.Store(test1); err != nil {
// 		panic(err)
// 	}

// 	test2 := cargo.New("ABC123", cargo.RouteSpecification{
// 		Origin:          location.SESTO,
// 		Destination:     location.CNHKG,
// 		ArrivalDeadline: time.Now().AddDate(0, 0, 14),
// 	})
// 	if err := r.Store(test2); err != nil {
// 		panic(err)
// 	}
// }
