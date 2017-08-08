package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/geneseeq/authorize-system/cms/association/groups"
	"github.com/geneseeq/authorize-system/cms/association/users"
	"github.com/geneseeq/authorize-system/cms/grouping"
	"github.com/geneseeq/authorize-system/cms/roleing"
	"github.com/geneseeq/authorize-system/cms/route"
	"github.com/geneseeq/authorize-system/cms/usering"

	"github.com/go-kit/kit/log"
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

	fieldKeys := []string{"method"}

	gs := route.InitGroupRouter(logger, fieldKeys)
	as := route.InitRelationRouter(logger, fieldKeys)
	us := route.InitUserRouter(logger, fieldKeys)
	rs := route.InitRoleRouter(logger, fieldKeys)
	gus := route.InitUserRelationRouter(logger, fieldKeys)
	grs := route.InitRoleRelationRouter(logger, fieldKeys)

	httpLogger := log.With(logger, "component", "http")

	mux := http.NewServeMux()

	mux.Handle("/grouping/v1/", grouping.MakeHandler(gs, httpLogger))
	mux.Handle("/usering/v1/", usering.MakeHandler(us, httpLogger))
	mux.Handle("/roleing/v1/", roleing.MakeHandler(rs, httpLogger))
	mux.Handle("/releation/v1/user/", users.MakeHandler(as, httpLogger))
	mux.Handle("/releation/v1/group/", groups.MakeHandler(gus, grs, httpLogger))
	// mux.Handle("/gys/v1/", groups.MakeRoleHandler(grs, httpLogger))

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
