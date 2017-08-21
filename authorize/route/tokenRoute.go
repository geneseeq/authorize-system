package route

import (
	"github.com/geneseeq/authorize-system/authorize/action"
	"github.com/geneseeq/authorize-system/authorize/authing"

	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func initTokenRouter(logger log.Logger, fieldKeys []string) authing.Service {
	var datas = action.NewTokenDBRepository("test", "auth_infos")
	var rs authing.Service
	rs = authing.NewService(datas)
	rs = authing.NewLoggingService(log.With(logger, "component", "authing"), rs)
	rs = authing.NewInstrumentingService(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "authing_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "authing_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		rs,
	)
	return rs
}
