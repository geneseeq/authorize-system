package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	mgo "gopkg.in/mgo.v2"

	"github.com/geneseeq/authorize-system/cms/action"
	"github.com/geneseeq/authorize-system/cms/usering"

	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gopkg.in/yaml.v2"
)

const (
	defaultPort              = "8080"
	defaultRoutingServiceURL = "http://localhost:7878"
)

type Mongo struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type DB struct {
	Mongo Mongo `yaml:"mongo"`
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func initCfg() (string, error) {
	content, _ := ioutil.ReadFile("conf.yaml")
	mongoCfg := DB{}
	err := yaml.Unmarshal(content, &mongoCfg)
	if err != nil {
		return "", err
	}
	address := strings.Join([]string{
		mongoCfg.Mongo.Host, ":", mongoCfg.Mongo.Port,
	}, "")
	return address, nil
}

func main() {
	var (
		addr     = envString("PORT", defaultPort)
		httpAddr = flag.String("http.addr", ":"+addr, "HTTP listen address")
	)

	flag.Parse()

	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	address, err := initCfg()
	checkErr(err)
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:   []string{address},
		Timeout: 60 * time.Second,
	})
	checkErr(err)
	defer session.Close()

	collection := session.DB("test").C("user")

	var (
		users = action.NewUserDBRepository(collection)
	)

	fieldKeys := []string{"method"}

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

	mux.Handle("/usering/v1/", usering.MakeHandler(bs, httpLogger))

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
